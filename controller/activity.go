package controller

import (
	"BelajarAPI/controller/message"
	"BelajarAPI/controller/tools"
	"BelajarAPI/middlewares"
	"BelajarAPI/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (us *UserController) Add_Activity() echo.HandlerFunc {
	return func(c echo.Context) error {
		var activity model.Activity
		err := c.Bind(&activity)
		if err != nil {
			return c.JSON(message.Response(http.StatusBadRequest, "An error occurred while reading the input"))
		} else if activity.Do == "" {
			return c.JSON(message.Response(http.StatusBadRequest, "'DO' cannot be empty"))
		}

		activity.Email = middlewares.DecodeToken(c.Get("user"))

		isNotConflict, err := us.Model.Add_Activity(&activity)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while checking the activity"))
		} else if !isNotConflict {
			return c.JSON(message.Response(http.StatusConflict, "There are already activities at that time"))
		} else {
			return c.JSON(message.Response(http.StatusOK, "Succeeded in adding activity", map[string]any{"activity": activity}))
		}
	}
}

func (us *UserController) Edit_Activity() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err_id := strconv.Atoi(c.Param("id"))
		if err_id != nil {
			return c.JSON(message.Response(http.StatusBadRequest, "An error occurred while reading the ID"))
		}

		var activity model.Activity
		err_bind := c.Bind(&activity)
		if err_bind != nil {
			return c.JSON(message.Response(http.StatusBadRequest, "An error occurred while reading the input"))
		} else if activity.Do == "" {
			return c.JSON(message.Response(http.StatusBadRequest, "'DO' cannot be empty"))
		}

		activity.Email = middlewares.DecodeToken(c.Get("user"))

		err := us.Model.Edit_Activity(id, &activity)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while checking the activity"))
		} else {
			return c.JSON(message.Response(http.StatusOK, "Succeeded in update activity"))
		}
	}
}

func (us *UserController) Get_Activities() echo.HandlerFunc {
	return func(c echo.Context) error {
		activities, err := us.Model.Get_Activities(middlewares.DecodeToken(c.Get("user")))
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while checking the activity"))
		} else if len(activities) == 0 {
			return c.JSON(message.Response(http.StatusNotFound, "No activity at all"))
		}

		if activityResponse, err := tools.TypeConverter[[]message.ActivityResponse](&activities); err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while converting the activity"))
		} else {
			return c.JSON(message.Response(http.StatusOK, "Succeeded get all activities", map[string]any{"activity": activityResponse}))
		}
	}
}
