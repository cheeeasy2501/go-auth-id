package app

import (
	"context"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"user_id"`
}

type IAuthorizationService interface {
	LoginByEmail(email, password string) string
	Logout()
	
	GenerateToken(user *UserEntity) (string, error)
	ParseToken(accessToken string) (int64, error)
	HashPassword(password string) ([]byte, error)
	VerifyPassword(userPass, credentialsPass string) error
}

type AuthorizationService struct {
	secretKey string
}

func NewAuthorizationService(secretKey string) *AuthorizationService {
	return &AuthorizationService{
		secretKey: secretKey,
	}
}

func (auth *AuthorizationService) LoginByEmail(email, password string) string {
	return ""
}

func (auth *AuthorizationService) Logout() {
	
}


func (auth *AuthorizationService) GenerateToken(usr *UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: usr.Id,
	})

	return token.SignedString([]byte(auth.secretKey))
}

func (auth *AuthorizationService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid method")
		}
		return []byte(auth.secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Unathorized")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, err
	}

	return claims.UserId, nil
}


func (auth *AuthorizationService) HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return bytes, err
}

func (auth *AuthorizationService) VerifyPassword(userPass, credentialsPass string) error {
	comparePass, err := base64.StdEncoding.DecodeString(userPass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(comparePass, []byte(credentialsPass))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New("Invalid credentionals!")
		}

		return err
	}

	return nil
}