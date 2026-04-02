package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/superdima3000/transport-auth/db"
	"github.com/superdima3000/transport-auth/pkg/repository"
)

type AuthService struct {
	repo *repository.Repository
}

const tokenTTL = time.Hour * 12

type TokenClaims struct {
	jwt.StandardClaims
	UserID  int64 `json:"user_id"`
	IsAdmin int64 `json:"is_admin" db:"is_admin"`
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user db.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(login string, password string) (string, error) {
	user, err := s.repo.GetUserByUsernameAndPassword(login, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix()},
		UserID:  user.Id,
		IsAdmin: user.IsAdmin,
	})

	return token.SignedString([]byte(viper.GetString("jwt.key")))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(viper.GetString("crypto.salt"))))
}

func (s *AuthService) ParseToken(token string) (*TokenClaims, error) {
	tokenJWT, err := jwt.ParseWithClaims(token, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("jwt.key")), nil
	})

	if err != nil {
		return &TokenClaims{}, err
	}

	claims, ok := tokenJWT.Claims.(*TokenClaims)

	if !ok {
		return &TokenClaims{}, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
