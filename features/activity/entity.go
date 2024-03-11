package activity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ActivityController interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type ActivityService interface {
	AddActivity(token *jwt.Token, new_activity Activity) error
	UpdateActivity(token *jwt.Token, id_activity string, new_activity Activity) error
	DeleteActivity(token *jwt.Token, ActivityID string) error
	GetActivities(token *jwt.Token) ([]Activity, error)
}

type ActivityModul interface {
	AddActivity(email string, a Activity) error
	UpdateActivity(email string, id int, activity string) error
	DeleteActivity(email string, ActivityID string) error
	GetActivitiesByOwner(email string) ([]Activity, error)
}

type Activity struct {
	gorm.Model
	Email    string
	Activity string
}
