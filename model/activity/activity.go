package activity

import (
	"strings"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Email string
	Do    string `json:"do" form:"do"`
}

type ActivityModel struct {
	Connection *gorm.DB
	Activity   Activity
}

func (m *ActivityModel) Add_Activity(activity *Activity) (bool, error) {
	err := m.Connection.Create(&activity).Error
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

func (m *ActivityModel) Edit_Activity(id int, activity *Activity) error {
	query := m.Connection.Where("id = ?", uint(id)).Model(&activity).Updates(&activity)
	if err := query.Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (m *ActivityModel) Get_Activities(email string) ([]Activity, error) {
	var activities []Activity
	err := m.Connection.Where("email = ?", email).Find(&activities).Error
	if err != nil {
		return []Activity{}, err
	} else {
		return activities, nil
	}
}