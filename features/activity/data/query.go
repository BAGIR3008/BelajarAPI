package data

import (
	"BelajarAPI/features/activity"
	"BelajarAPI/helper"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) activity.ActivityModul {
	return &model{
		connection: db,
	}
}

func (m *model) AddActivity(email string, a activity.Activity) error {
	return m.connection.Create(&Activity{Email: email, Activity: a.Activity}).Error
}

func (m *model) UpdateActivity(email string, id int, activity string) error {
	if query := m.connection.Model(&Activity{}).Where("email = ? AND id = ?", email, id).Update("Activity", activity); query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}

func (m *model) GetActivitiesByOwner(email string) ([]activity.Activity, error) {
	var activities []activity.Activity
	if err := m.connection.Where("email = ?", email).Find(&activities).Error; err != nil {
		return []activity.Activity{}, err
	}
	return activities, nil
}

func (m *model) DeleteActivity(email string, ActivityID string) error {
	if query := m.connection.Where("email = ? AND id = ?", email, ActivityID).Delete(&activity.Activity{}); query.Error != nil {
		return query.Error
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}
