package app

import (
	"gorm.io/gorm"
)

type UserEntity struct {
	gorm.Model
	FirstName, LastName, MiddleName, password string
	isBanned bool
}


func NewUserEntity() UserEntity {
	return UserEntity{}
}
