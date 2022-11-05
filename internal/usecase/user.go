package usecase

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"
	"weather_api/config"
	postgressql "weather_api/internal/adapters/db/postgres_sql"
	"weather_api/internal/entities"

	"github.com/golang-jwt/jwt"
)

type authorizationUseCase struct {
	storage postgressql.AuthorizationStorage
}

func NewAuthorizationUseCase(storage postgressql.AuthorizationStorage) *authorizationUseCase {
	return &authorizationUseCase{
		storage: storage,
	}
}

func (u *authorizationUseCase) CreateUser(user entities.User) (id int, err error) {
	user.Password = generatePasswordHash(user.Password)
	return u.storage.CreateUser(user)
}

func (u *authorizationUseCase) GenerateToken(login, password string) (string, error) {
	user, err := u.storage.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &entities.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(config.SignedKey))
}

func (u *authorizationUseCase) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &entities.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(config.SignedKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*entities.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *entities.TokenClaims")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(config.Salt)))
}
