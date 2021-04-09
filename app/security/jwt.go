package security

import (
	"github.com/google/uuid"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app/config"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/go-logger/v2"
)

type JWT struct {
	jwtSignKey []byte
}

// New инициализирует сервер для ответа на сетевые запросы HTTP.
func New(cfg *config.Config) *JWT {
	return &JWT{jwtSignKey: []byte(cfg.Server.JWTSignKey)}
}

type Session struct {
	jwt.StandardClaims
}

// GenerateToken создаёт токен.
func (j *JWT) GenerateToken(sessionId uuid.UUID) (string, time.Time) {

	if logger.DebugEnabled() {
		logger.Debugf("Create new session %s", sessionId)
	}

	expireAt := time.Now().Add(1 * time.Minute)

	// Вставить информацию о сессии пользователя в `token`
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Session{
		StandardClaims: jwt.StandardClaims{
			Id:        sessionId.String(),
			ExpiresAt: expireAt.Unix(),
		},
	})

	// token -> string.
	tokenString, err := newToken.SignedString(j.jwtSignKey)
	if err != nil {
		logger.Error(err)
	}

	return tokenString, expireAt
}

// ValidateToken проверяет токен.
func (j *JWT) ValidateToken(requestToken string) (*jwt.Token, error) {

	if logger.DebugEnabled() {
		logger.Debug("Validating token...")
	}
	session := &Session{}
	token, err := jwt.ParseWithClaims(requestToken, session, func(token *jwt.Token) (interface{}, error) {
		return j.jwtSignKey, nil
	})

	return token, err
}
