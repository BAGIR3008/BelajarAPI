package user

import (
	"BelajarAPI/model/activity"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string              `json:"name" form:"name"`
	Email      string              `json:"email" form:"email" gorm:"unique"`
	Password   string              `json:"password" form:"password"`
	Activities []activity.Activity `gorm:"foreignKey:Email;references:Email"`
}

type UserModel struct {
	Connection *gorm.DB
	User       User
}

func (um *UserModel) CheckByID(id int) bool {
	var user User
	err := um.Connection.Where("id = ?", uint(id)).Find(&user).Error
	if err != nil || user.Name == "" {
		return false
	}

	return true
}

func (um *UserModel) CheckByEmail(email string) bool {
	var user User
	err := um.Connection.Find(&user, User{Email: email}).Error
	if err != nil || user.Email == "" {
		return false
	}

	return true
}

func (um *UserModel) GetUsers() ([]User, error) {
	var users []User
	err := um.Connection.Find(&users).Error
	if err != nil {
		return []User{}, err
	} else {
		return users, nil
	}
}

func (um *UserModel) GetUserByID(id int) (User, error) {
	var user User
	err := um.Connection.Where("id = ?", uint(id)).Find(&user).Error
	if err != nil {
		return User{}, err
	} else if user.Name == "" {
		return User{}, nil
	} else {
		return user, nil
	}
}

func (um *UserModel) DeleteUserByID(id int) (bool, error) {
	var user User
	query := um.Connection.Delete(&user, id)
	if err := query.Error; err != nil {
		return false, err
	} else if !(query.RowsAffected > 0) {
		return false, nil
	} else {
		return true, nil
	}
}

func (um *UserModel) UpdateUserByID(id int, user User) error {
	query := um.Connection.Where("id = ?", uint(id)).Model(&user).Updates(&user)
	if err := query.Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (um *UserModel) Login(email string, password string) (User, error) {
	var result User
	err := um.Connection.Where("email = ? AND password = ?", email, password).Find(&result).Error
	if err != nil {
		return User{}, err
	} else if result.Name == "" {
		return User{}, nil
	} else {
		return result, nil
	}
}

func (um *UserModel) Register(newUser *User) (bool, error) {
	err := um.Connection.Create(&newUser).Error
	if err != nil {
		if contain := strings.Contains(err.Error(), "Duplicate entry"); contain {
			return false, nil
		} else {
			return false, err
		}
	} else {
		return true, nil
	}
}

func (um *UserModel) Profile(email string) (User, error) {
	var result User
	err := um.Connection.Where("email = ?", email).Find(&result).Error
	if err != nil {
		return User{}, err
	} else if result.Name == "" {
		return User{}, nil
	} else {
		return result, nil
	}
}
