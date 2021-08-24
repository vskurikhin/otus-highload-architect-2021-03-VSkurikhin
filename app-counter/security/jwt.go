package security

import (
	"errors"
	"github.com/vskurikhin/otus-highload-architect-2021-03-VSkurikhin/app-counter/config"
	"strconv"
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
func (j *JWT) GenerateToken(sessionId uint64) (string, time.Time) {

	if logger.DebugEnabled() {
		logger.Debugf("UpdateOrCreate new session %s", sessionId)
	}
	expireAt := time.Now().Add(1000 * time.Minute)

	// Вставить информацию о сессии пользователя в `token`
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &Session{
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.FormatUint(sessionId, 10),
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

// SessionIdFromToken получить из токена id-сессии.
func (j *JWT) SessionIdFromToken(sToken string) (*string, error) {

	token, err := jwt.ParseWithClaims(sToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.jwtSignKey, nil
	})

	if err != nil {
		return nil, err // правильная обработка ошибок вместо паники
	}
	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		if logger.DebugEnabled() {
			logger.Debugf("%v %v", claims.Id, claims.ExpiresAt)
		}
		return &claims.Id, nil
	}
	return nil, errors.New("")
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
