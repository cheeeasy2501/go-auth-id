package app

import (
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	FirstName, LastName, MiddleName, password, Email string
	isBanned                                         bool
}

func NewUserEntity() UserEntity {
	return UserEntity{}
}

func (u *UserEntity) Password() string {
	return u.password
}

func (u *UserEntity) SetPassword(password string) {
	u.password = password
}
