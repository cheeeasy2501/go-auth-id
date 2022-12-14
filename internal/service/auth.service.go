package service

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/cheeeasy2501/auth-id/internal/entity"
	"github.com/cheeeasy2501/auth-id/internal/transport/http/v1/request"
)

type UserClaims struct {
	Id         uint   `json:"id"`
	Avatar     string `json:"avatar"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserClaims `json:"user"`
}

type IAuthorizationService interface {
	Registration(request *request.RegistrationRequest) error
	LoginByEmail(request *request.LoginByEmailRequest) (string, error)
	Logout()

	GenerateToken(user *entity.User) (string, error)
	ParseToken(accessToken string) (uint, error)
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

func (s *AuthorizationService) LoginByEmail(request *request.LoginByEmailRequest) (string, error) {
	user := entity.NewUser()
	tx := s.conn.First(&user, "email = $1", request.Email)

	if tx.RowsAffected == 0 {
		return "", errors.New("Login or password not found")
	}

	if err := s.VerifyPassword(user.Password, request.Password); err != nil {
		return "", err
	}

	token, err := s.GenerateToken(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthorizationService) Logout() {

}

func (s *AuthorizationService) GenerateToken(usr *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserClaims: UserClaims{
			Id:         usr.ID,
			Avatar:     usr.Avatar,
			FirstName:  usr.FirstName,
			LastName:   usr.LastName,
			MiddleName: usr.MiddleName,
		},
	})

	return token.SignedString([]byte(s.secretKey))
}

func (s *AuthorizationService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
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

func (s *AuthorizationService) HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return bytes, err
}

func (s *AuthorizationService) VerifyPassword(userPass, credentialsPass string) error {
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