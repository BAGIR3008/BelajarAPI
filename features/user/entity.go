package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface {
	Login() echo.HandlerFunc
	Register() echo.HandlerFunc
	// Update() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token) (User, error)
}

type UserModel interface {
	InsertUser(newData User) error
	UpdateUser(email string, data User) error
	Login(hp string) (User, error)
	GetUserByEmail(email string) (User, error)
}

type User struct {
	gorm.Model
	Name     string
	Hp       string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanumunicode"`
}

type Register struct {
	Name     string `validate:"required,alpha"`
	Hp       string `validate:"required,number"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanumunicode"`
}
