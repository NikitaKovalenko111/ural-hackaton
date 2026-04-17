package auth_service

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	auth_dto "ural-hackaton/internal/dto/auth"
	"ural-hackaton/internal/services/handlers/email" // ваш пакет отправки писем

	// репозитории
	auth_storage "ural-hackaton/internal/storage/repositories/auth"
	users_storage "ural-hackaton/internal/storage/repositories/user"
	// модели пользователей
)

type AuthService struct {
	userRepo  *users_storage.UserRepo
	tokenRepo *auth_storage.AuthTokenRepo
	emailSrv  *email.EmailSender
	// Для генерации сессий (упрощённо, лучше использовать JWT библиотеку)
	secretKey         []byte
	allowDevMagicLink bool
}

const sessionTTL = 24 * time.Hour

func Init(
	userRepo *users_storage.UserRepo,
	tokenRepo *auth_storage.AuthTokenRepo,
	emailSrv *email.EmailSender,
	secretKey []byte,
	allowDevMagicLink bool,
) *AuthService {
	return &AuthService{
		userRepo:          userRepo,
		tokenRepo:         tokenRepo,
		emailSrv:          emailSrv,
		secretKey:         secretKey,
		allowDevMagicLink: allowDevMagicLink,
	}
}

// RequestMagicLink: логика запроса ссылки
func (s *AuthService) RequestMagicLink(email string) error {
	// 1. Ищем пользователя
	usr, err := s.userRepo.GetUserByEmail(email)
	if errors.Is(err, sql.ErrNoRows) {
		// Не раскрываем, что пользователя нет (защита от перебора)
		// Просто возвращаем nil, контроллер скажет "Отправлено"
		return nil
	}
	if err != nil {
		return fmt.Errorf("get user by email: %w", err)
	}

	// 2. Генерируем токен
	token, tokenHash, err := generateSecureToken()
	if err != nil {
		return fmt.Errorf("generate token: %w", err)
	}

	// 3. Сохраняем хеш в БД (срок жизни 15 минут)
	// Предполагается, что в AuthTokenRepo есть метод SaveTokenHash
	err = s.tokenRepo.SaveTokenHash(usr.Id, email, tokenHash, 15*time.Minute)
	if err != nil {
		return fmt.Errorf("save token: %w", err)
	}

	// 4. Отправляем письмо синхронно, чтобы вернуть клиенту реальную ошибку SMTP,
	// если доставка не удалась.
	if err := s.emailSrv.SendMagicLink(usr.Email, token); err != nil {
		return fmt.Errorf("send magic link: %w", err)
	}

	return nil
}

// RequestMagicLinkDev: только для локальной разработки.
// Возвращает magic-link напрямую в ответе вместо отправки письма.
func (s *AuthService) RequestMagicLinkDev(email string) (string, error) {
	if !s.allowDevMagicLink {
		return "", errors.New("dev magic link is disabled")
	}

	usr, err := s.userRepo.GetUserByEmail(email)
	if errors.Is(err, sql.ErrNoRows) {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", fmt.Errorf("get user by email: %w", err)
	}

	token, tokenHash, err := generateSecureToken()
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	if err := s.tokenRepo.SaveTokenHash(usr.Id, email, tokenHash, 15*time.Minute); err != nil {
		return "", fmt.Errorf("save token: %w", err)
	}

	link := fmt.Sprintf("%s/auth/verify?token=%s", s.emailSrv.FrontendURL, token)
	return link, nil
}

// VerifyMagicLink: проверка токена и вход
func (s *AuthService) VerifyMagicLink(token string) (*auth_dto.VerifyMagicLinkResponse, error) {
	// 1. Валидируем токен через репозиторий
	// Метод должен: найти хеш, проверить срок, проверить used_at, пометить как использованный
	userID, err := s.tokenRepo.ValidateAndConsumeToken(token)
	if err != nil {
		return nil, err // Ошибки: invalid, expired, already_used
	}

	// 2. Загружаем данные пользователя
	usr, err := s.userRepo.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("get user after verify: %w", err)
	}

	// 3. Генерируем сессию (или JWT)
	// Здесь упрощённая генерация, в продакшене используйте github.com/golang-jwt/jwt/v5
	sessionToken, err := generateSessionToken(userID, s.secretKey, sessionTTL)
	if err != nil {
		return nil, fmt.Errorf("generate session: %w", err)
	}

	return &auth_dto.VerifyMagicLinkResponse{
		UserID:       usr.Id,
		Fullname:     usr.FullName,
		Email:        usr.Email,
		Role:         usr.Role,
		Telegram:     usr.Telegram,
		Phone:        usr.Phone,
		SessionToken: sessionToken,
	}, nil
}

func (s *AuthService) GetSessionUser(sessionToken string) (*auth_dto.VerifyMagicLinkResponse, error) {
	if sessionToken == "" {
		return nil, fmt.Errorf("session token is empty")
	}

	userID, err := parseAndValidateSessionToken(sessionToken, s.secretKey)
	if err != nil {
		return nil, fmt.Errorf("token validation failed: %w", err)
	}

	usr, err := s.userRepo.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("get user by session: %w", err)
	}

	return &auth_dto.VerifyMagicLinkResponse{
		UserID:   usr.Id,
		Fullname: usr.FullName,
		Email:    usr.Email,
		Role:     usr.Role,
		Telegram: usr.Telegram,
		Phone:    usr.Phone,
	}, nil
}

// --- Вспомогательные функции ---

func generateSecureToken() (string, string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", err
	}

	token := base64.URLEncoding.EncodeToString(b)
	hash := sha256.Sum256([]byte(token))
	tokenHash := base64.URLEncoding.EncodeToString(hash[:])

	return token, tokenHash, nil
}

func generateSessionToken(userID uint64, secret []byte, ttl time.Duration) (string, error) {
	expiresAt := time.Now().Add(ttl).Unix()
	payload := fmt.Sprintf("%d:%d", userID, expiresAt)

	mac := hmac.New(sha256.New, secret)
	if _, err := mac.Write([]byte(payload)); err != nil {
		return "", err
	}

	signature := mac.Sum(nil)
	payloadPart := base64.RawURLEncoding.EncodeToString([]byte(payload))
	signaturePart := base64.RawURLEncoding.EncodeToString(signature)

	return payloadPart + "." + signaturePart, nil
}

func parseAndValidateSessionToken(token string, secret []byte) (uint64, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return 0, errors.New("invalid session token format")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return 0, errors.New("invalid session token payload")
	}

	signatureBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, errors.New("invalid session token signature")
	}

	mac := hmac.New(sha256.New, secret)
	if _, err := mac.Write(payloadBytes); err != nil {
		return 0, err
	}

	expectedSignature := mac.Sum(nil)
	if !hmac.Equal(signatureBytes, expectedSignature) {
		return 0, errors.New("invalid session token signature")
	}

	payloadParts := strings.Split(string(payloadBytes), ":")
	if len(payloadParts) != 2 {
		return 0, errors.New("invalid session token payload content")
	}

	userID, err := strconv.ParseUint(payloadParts[0], 10, 64)
	if err != nil {
		return 0, errors.New("invalid session token user id")
	}

	expiresAtUnix, err := strconv.ParseInt(payloadParts[1], 10, 64)
	if err != nil {
		return 0, errors.New("invalid session token expiration")
	}

	if time.Now().After(time.Unix(expiresAtUnix, 0)) {
		return 0, errors.New("session token expired")
	}

	return userID, nil
}
