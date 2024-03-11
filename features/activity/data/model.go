package data

import "gorm.io/gorm"

type Activity struct {
	gorm.Model
	Email    string
	Activity string
}
