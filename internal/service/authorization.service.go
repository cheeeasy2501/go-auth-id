package service

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/cheeeasy2501/auth-id/internal/apperr"
	"github.com/cheeeasy2501/auth-id/internal/entity"
	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/request"
)

type UserClaims struct {
	Id uint64 `json:"id"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserClaims `json:"user"`
}

type RefreshClaims struct {
	jwt.RegisteredClaims
	Id uint64 `json:"id"`
}

type ITokenService interface {
	generateAccessToken(user *entity.User) (string, error)
	generateRefreshToken(userId uint64) (string, error)
	ParseToken(t string) (uint64, error)
	ParseRefreshToken(t string) (uint64, error)
	RefreshToken(request *request.RefreshTokenRequest) (entity.Tokens, error)
}

type IAuthorizationService interface {
	ITokenService
	Registration(request *request.RegistrationRequest) error
	LoginByEmail(request *request.LoginByEmailRequest) (entity.Tokens, error)
	Logout()
	HashPassword(password string) ([]byte, error)
	VerifyPassword(userPass, credentialsPass string) error
}

type AuthorizationService struct {
	conn      *gorm.DB
	secretKey string
}

func NewAuthorizationService(secretKey string, conn *gorm.DB) *AuthorizationService {
	return &AuthorizationService{
		secretKey: secretKey,
		conn:      conn,
	}
}

func (s *AuthorizationService) Registration(request *request.RegistrationRequest) error {

	hashPassword, err := s.HashPassword(request.Password)
	if err != nil {
		return err
	}

	encryptedPass := base64.StdEncoding.EncodeToString(hashPassword)
	request.Password = encryptedPass

	user := entity.NewUserFromRegistrationRequest(*request)

	result := s.conn.First(&user, "email = ?", user.Email)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return err
	}

	if result.RowsAffected > 0 {
		return errors.New("User already exist")
	}

	err = s.conn.Create(&user).Error
	if err != nil {
		return errors.New("User not created")
	}

	return nil
}

func (s *AuthorizationService) LoginByEmail(request *request.LoginByEmailRequest) (entity.Tokens, error) {
	user := entity.NewUser()
	tx := s.conn.First(&user, "email = $1", request.Email)

	if tx.RowsAffected == 0 {
		return entity.Tokens{}, errors.New("Login or password not found")
	}

	if err := s.VerifyPassword(user.Password, request.Password); err != nil {
		return entity.Tokens{}, err
	}

	accessToken, err := s.generateAccessToken(&user)
	if err != nil {
		return entity.Tokens{}, err
	}

	refreshToken, err := s.generateRefreshToken(user.ID)
	if err != nil {
		return entity.Tokens{}, err
	}

	tokens := entity.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (s *AuthorizationService) Logout() {

}

// Генерируем access token
func (s *AuthorizationService) generateAccessToken(usr *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserClaims: UserClaims{
			Id: usr.ID,
		},
	})

	return token.SignedString([]byte(s.secretKey))
}

// Генерируем refresh token
func (s *AuthorizationService) generateRefreshToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Id: userId,
	})

	return token.SignedString([]byte(s.secretKey))
}

// Парсим и валидируем токен
func (s *AuthorizationService) ParseToken(t string) (uint64, error) {
	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Unathorized")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, err
	}

	return claims.Id, nil
}

func (s *AuthorizationService) ParseRefreshToken(t string) (uint64, error) {
	token, err := jwt.ParseWithClaims(t, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid method")
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Unathorized")
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok {
		return 0, err
	}

	return claims.Id, nil
}

// Хэшируем пароль в bcrypt
func (s *AuthorizationService) HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return bytes, err
}

// Проверка пароля
func (s *AuthorizationService) VerifyPassword(userPass, credentialsPass string) error {
	comparePass, err := base64.StdEncoding.DecodeString(userPass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(comparePass, []byte(credentialsPass))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return new(apperr.InvalidCredentionals)
		}

		return err
	}

	return nil
}

// Возвращаем обновленные токены
func (s *AuthorizationService) RefreshToken(request *request.RefreshTokenRequest) (entity.Tokens, error) {
	newTokens := new(entity.Tokens)
	usr := new(entity.User)
	tx := s.conn.First(usr, "id = ?", request.UserId)

	if tx.RowsAffected == 0 {
		return *newTokens, errors.New("Invalid request") // TODO: to apperrors
	}

	accessToken, err := s.generateAccessToken(usr)
	if err != nil {
		return *newTokens, errors.New("Invalid request") // TODO: to apperrors
	}

	refreshToken, err := s.generateRefreshToken(usr.ID)
	if err != nil {
		return *newTokens, errors.New("Invalid request") // TODO: to apperrors
	}

	newTokens.AccessToken = accessToken
	newTokens.RefreshToken = refreshToken

	return *newTokens, nil
}
