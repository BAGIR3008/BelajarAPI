package data

import (
	"BelajarAPI/features/user"
	"BelajarAPI/helper"
	"errors"
	"log"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) user.UserModel {
	return &model{
		connection: db,
	}
}

func (m *model) InsertUser(newData user.User) error {
	return m.connection.Create(&newData).Error
}

func (m *model) UpdateUser(email string, data user.User) error {
	if query := m.connection.Model(&data).Where("email = ?", email).Omit("email").Updates(&data); query.Error != nil {
		log.Println(query.Error.Error())
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorNoRowsAffected)
	}
	return nil
}

func (m *model) GetUserByEmail(email string) (user.User, error) {
	var result user.User
	if err := m.connection.Where("email = ?", email).Find(&result).Error; err != nil {
		log.Println(err.Error())
		return user.User{}, errors.New(helper.ErrorGeneralDatabase)
	} else if result.Email == "" {
		return user.User{}, errors.New(helper.ErrorDatabaseNotFound)
	}
	return result, nil
}

func (m *model) Login(email string) (user.User, error) {
	var result user.User
	if err := m.connection.Where("email = ? ", email).First(&result).Error; err != nil {
		if result.Email == "" {
			return user.User{}, nil
		} else {
			return user.User{}, errors.New(helper.ErrorGeneralDatabase)
		}
	}
	return result, nil
}
