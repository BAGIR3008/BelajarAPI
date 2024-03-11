package services

import (
	"BelajarAPI/features/user"
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	model user.UserModel
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.UserModel) user.UserService {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(newData user.User) error {
	var registerValidate user.Register
	registerValidate.Name = newData.Name
	registerValidate.Email = newData.Email
	registerValidate.Hp = newData.Hp
	registerValidate.Password = newData.Hp
	err := s.v.Struct(&registerValidate)
	if err != nil {
		log.Println("error validate", err.Error())
		return err
	}

	newPassword, err := s.pm.HashPassword(newData.Password)
	if err != nil {
		return errors.New(helper.ErrorGeneralServer)
	}
	newData.Password = newPassword

	err = s.model.InsertUser(newData)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return errors.New(mysqlErr.Message)
		}
		return errors.New(helper.ErrorGeneralServer)
	}

	return nil
}

func (s *service) Login(loginData user.User) (user.User, string, error) {
	var loginValidate user.Login
	loginValidate.Email = loginData.Email
	loginValidate.Password = loginData.Password
	err := s.v.Struct(&loginValidate)
	if err != nil {
		log.Println("error validate", err.Error())
		return user.User{}, "", err
	}

	dbData, err := s.model.Login(loginValidate.Email)
	if err != nil {
		return user.User{}, "", err
	}

	if err := s.pm.ComparePassword(loginData.Password, dbData.Password); err != nil {
		return user.User{}, "", errors.New(helper.ErrorUserCredential)
	}

	token, err := middlewares.GenerateJWT(dbData.Email)
	if err != nil {
		return user.User{}, "", errors.New(helper.ErrorGeneralServer)
	}

	return dbData, token, nil
}

func (s *service) Profile(token *jwt.Token) (user.User, error) {
	decodeEmail := middlewares.DecodeToken(token)
	result, err := s.model.GetUserByEmail(decodeEmail)
	if err != nil {
		return user.User{}, err
	}

	return result, nil
}
