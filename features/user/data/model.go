package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Hp       string
	Email    string
	Password string
	// Activities []data.Activity
}
