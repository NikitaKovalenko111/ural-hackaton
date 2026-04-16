package auth_service

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
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
	secretKey []byte
}

func Init(
	userRepo *users_storage.UserRepo,
	tokenRepo *auth_storage.AuthTokenRepo,
	emailSrv *email.EmailSender,
	secretKey []byte,
) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		emailSrv:  emailSrv,
		secretKey: secretKey,
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
	sessionToken, err := generateSessionToken(userID, s.secretKey)
	if err != nil {
		return nil, fmt.Errorf("generate session: %w", err)
	}

	return &auth_dto.VerifyMagicLinkResponse{
		UserID:       usr.Id,
		Fullname:     usr.FullName,
		Email:        usr.Email,
		Role:         usr.Role,
		SessionToken: sessionToken,
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

func generateSessionToken(userID uint64, secret []byte) (string, error) {
	// Упрощённый пример: хеш от ID + время + секрет
	// В реальности: JWT с экспайром, issuer, subject и т.д.
	raw := fmt.Sprintf("%d:%d:%s", userID, time.Now().UnixNano(), string(secret))
	h := sha256.Sum256([]byte(raw))
	return base64.URLEncoding.EncodeToString(h[:]), nil
}
