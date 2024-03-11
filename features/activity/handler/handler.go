package handler

import (
	"BelajarAPI/features/activity"
	"BelajarAPI/helper"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s activity.ActivityService
}

func NewHandler(service activity.ActivityService) activity.ActivityController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ActivityResponse
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		var inputProcess activity.Activity
		inputProcess.Activity = input.Activity
		err = ct.s.AddActivity(token, inputProcess)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success add activity"))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ActivityResponse
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		var inputProcess activity.Activity
		inputProcess.Activity = input.Activity
		err = ct.s.UpdateActivity(token, c.Param("id"), inputProcess)
		if err != nil {
			if err.Error() == helper.ErrorDatabaseNotFound {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, err.Error()))
			}
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success change activity"))
	}
}

func (ct *controller) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		activities, err := ct.s.GetActivities(token)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, err.Error()))
		}

		var activitiesResponse []ActivityResponse
		if err := helper.Recast(&activities, &activitiesResponse); err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success get all activities", map[string]any{"activities": activitiesResponse}))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err := ct.s.DeleteActivity(token, c.Param("id"))
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success delete activity"))
	}
}
