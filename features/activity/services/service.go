package services

import (
	"BelajarAPI/features/activity"
	"BelajarAPI/helper"
	"BelajarAPI/middlewares"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	m activity.ActivityModul
	v *validator.Validate
}

func NewActivityService(model activity.ActivityModul) activity.ActivityService {
	return &service{
		m: model,
		v: validator.New(),
	}
}

func (s *service) AddActivity(token *jwt.Token, new_activity activity.Activity) error {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	if strings.TrimSpace(new_activity.Activity) == "" {
		return errors.New("activity cannot be empty")
	}

	err := s.v.Struct(&new_activity)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	err = s.m.AddActivity(email, new_activity)
	if err != nil {
		return errors.New(helper.ErrorGeneralServer)
	}

	return nil
}

func (s *service) UpdateActivity(token *jwt.Token, id_activity string, new_activity activity.Activity) error {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	id, err := strconv.Atoi(id_activity)
	if err != nil {
		return errors.New("invalid id")
	}

	if strings.TrimSpace(new_activity.Activity) == "" {
		return errors.New("activity cannot be empty")
	}

	err = s.v.Struct(&new_activity)
	if err != nil {
		log.Println("error validasi", err.Error())
		return err
	}

	err = s.m.UpdateActivity(email, id, new_activity.Activity)

	if err != nil {
		if err.Error() != helper.ErrorDatabaseNotFound {
			log.Println(err)
		}
		return err
	}

	return nil
}

func (s *service) GetActivities(token *jwt.Token) ([]activity.Activity, error) {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return []activity.Activity{}, errors.New("data tidak valid")
	}

	activities, err := s.m.GetActivitiesByOwner(email)
	if err != nil {
		return []activity.Activity{}, errors.New(helper.ErrorGeneralServer)
	}

	return activities, nil
}

func (s *service) DeleteActivity(token *jwt.Token, ActivityID string) error {
	email := middlewares.DecodeToken(token)
	if email == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.DeleteActivity(email, ActivityID)
	if err != nil {
		return errors.New(helper.ErrorDatabaseNotFound)
	}

	return nil
}
