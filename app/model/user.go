package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserDetails
	UserCredentials
}

type UserDetails struct {
	Username      string `gorm:"string;size:255;not null;unique;" json:"username"`
	Email         string `gorm:"string;size:100;unique;" json:"email"`
	EmailVerified bool   `gorm:"bool;not null;default:false" json:"emailVerified"`
	Firstname     string `gorm:"string;size:255;" json:"firstname"`
	Lastname      string `gorm:"string;size:255;" json:"lastname"`
}

type UserCredentials struct {
	Password string `gorm:"size:255;not null;" json:"-"`
}
