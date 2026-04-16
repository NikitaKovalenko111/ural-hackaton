package auth_repo

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"
	"ural-hackaton/internal/storage"
)

// internal/storage/auth_token_repo.go

type AuthTokenRepo struct {
	db *storage.Storage
}

func Init(db *storage.Storage) *AuthTokenRepo {
	return &AuthTokenRepo{
		db: db,
	}
}

// SaveTokenHash сохраняет хеш токена
func (r *AuthTokenRepo) SaveTokenHash(userID uint64, email, tokenHash string, ttl time.Duration) error {
	expiresAt := time.Now().Add(ttl)
	_, err := r.db.Db.Exec(`
		INSERT INTO auth_tokens (user_id, email, token_hash, expires_at)
		VALUES ($1, $2, $3, $4)
	`, userID, email, tokenHash, expiresAt)
	return err
}

// ValidateAndConsumeToken проверяет и "сжигает" токен
func (r *AuthTokenRepo) ValidateAndConsumeToken(token string) (uint64, error) {
	// Хешим входящий токен для поиска
	hash := sha256.Sum256([]byte(token))
	tokenHash := base64.URLEncoding.EncodeToString(hash[:])

	var userID uint64
	var expiresAt time.Time
	var usedAt sql.NullTime

	err := r.db.Db.QueryRow(`
		SELECT user_id, expires_at, used_at 
		FROM auth_tokens 
		WHERE token_hash = $1
	`, tokenHash).Scan(&userID, &expiresAt, &usedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, sql.ErrNoRows
	}
	if err != nil {
		return 0, err
	}

	// Проверка времени
	if time.Now().After(expiresAt) {
		return 0, errors.New("token expired")
	}
	// Проверка использования
	if usedAt.Valid {
		return 0, errors.New("token already used")
	}

	// Атомарно помечаем как использованный
	_, err = r.db.Db.Exec(`UPDATE auth_tokens SET used_at = NOW() WHERE token_hash = $1`, tokenHash)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
