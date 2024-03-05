package model

import (
	"BelajarAPI/model/activity"
	"BelajarAPI/model/user"

	"gorm.io/gorm"
)

func User(db *gorm.DB) *user.UserModel {
	return &user.UserModel{Connection: db}
}

func Activity(db *gorm.DB) *activity.ActivityModel {
	return &activity.ActivityModel{Connection: db}
}
