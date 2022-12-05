package app

import (
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	entity "github.com/cheeeasy2501/auth-id/internal/entity/app"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"user_id"`
}

type IAuthorizationService interface {
	// Registration(email, password)
	LoginByEmail(email, password string) (entity.UserEntity, error)
	Logout()

	GenerateToken(user *entity.UserEntity) (string, error)
	ParseToken(accessToken string) (uint, error)
	HashPassword(password string) ([]byte, error)
	VerifyPassword(userPass, credentialsPass string) error
}

type AuthorizationService struct {
	conn        *gorm.DB
	secretKey string
}

func NewAuthorizationService(secretKey string, conn *gorm.DB) *AuthorizationService {
	return &AuthorizationService{
		secretKey: secretKey,
		conn:        conn,
	}
}

func (s *AuthorizationService) Registration(user *entity.UserEntity) (string, error) {

	hashPassword, err := s.HashPassword(user.Password())
	if err != nil {
		return "", err
	}

	encryptedPass := base64.StdEncoding.EncodeToString(hashPassword)
	user.SetPassword(encryptedPass)

	tx := s.conn.First(user, "email = $1", user.Email)
	if tx.RowsAffected > 0 {
		return "", errors.New("User already exist")
	}

	tx = s.conn.Create(user)
	if tx.RowsAffected == 0 {
        return "", errors.New("User not created")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return  "", err
	}

	return token, nil
}

func (s *AuthorizationService) LoginByEmail(email, password string) (entity.UserEntity, error) {
	user := entity.NewUserEntity()
	tx := s.conn.First(&user, "email = $1", email)

	if tx.RowsAffected == 0 {
		return entity.UserEntity{}, errors.New("Login or password not found")
	}

	if err := s.VerifyPassword(user.Password(), password); err != nil {
		return entity.UserEntity{}, err
	}


	return user, nil
}

func (s *AuthorizationService) Logout() {

}

func (s *AuthorizationService) GenerateToken(usr *entity.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: usr.ID,
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

	return claims.UserId, nil
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
