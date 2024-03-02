package controller

import (
	"BelajarAPI/middlewares"
	"BelajarAPI/model"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

func response(status int, message string, data ...map[string]any) (int, map[string]any) {
	json := map[string]any{
		"code":    status,
		"message": message,
	}

	for _, part := range data {
		for key, value := range part {
			json[key] = value
		}
	}
	return status, json
}

// get users
func (us *UserController) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := us.Model.GetUsers()
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "An error occurred while calling user data"))
		} else if len(users) == 0 {
			return c.JSON(response(http.StatusNoContent, "Data is empty"))
		} else {
			return c.JSON(response(http.StatusOK, "success get all users", map[string]any{"users": users}))
		}
	}
}

// get user by id
func (us *UserController) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "An error occurred while reading the ID"))
		}

		user, err := us.Model.GetUserByID(id)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "An error occurred while calling user data"))
		} else if user.Name == "" {
			return c.JSON(response(http.StatusNotFound, "User ID Not Found"))
		} else {
			return c.JSON(response(http.StatusOK, "success get user", map[string]any{"user": user}))
		}
	}
}

// delete user by id
func (us *UserController) DeleteUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "An error occurred while reading the ID"))
		}

		success, err := us.Model.DeleteUserByID(id)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "An error occurred while checking the user"))
		} else if !success {
			return c.JSON(response(http.StatusNotFound, "User ID Not Found"))
		} else {
			return c.JSON(response(http.StatusOK, "success delete user"))
		}
	}
}

// update user by id
func (us *UserController) UpdateUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "An error occurred while reading the ID"))
		} else if !us.Model.CheckByID(id) {
			return c.JSON(response(http.StatusNotFound, "User ID Not Found"))
		}

		var user model.User
		err = c.Bind(&user)
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "An error occurred while reading the input"))
		}

		err = us.Model.UpdateUserByID(id, user)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "An error occurred while checking the user"))
			// } else if !changed {
			// 	return c.JSON(http.StatusBadRequest, map[string]any{
			// 		"code":     http.StatusBadRequest,
			// 		"messages": "No change occurred",
			// 	})
		} else {
			return c.JSON(response(http.StatusOK, "success update user"))
		}
	}
}

// Login
func (us *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userJSON model.User
		err := c.Bind(&userJSON)
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "An error occurred while reading the input"))
		}

		if userJSON.Email == "" || userJSON.Password == "" {
			return c.JSON(response(http.StatusBadRequest, "Email, and Password cannot be empty"))
		}

		user, err := us.Model.Login(userJSON.Email, userJSON.Password)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "An error occurred while checking the user"))
		}

		token, err := middlewares.GenerateJWT(user.Email)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "the crow makes tokens"))
		}

		return c.JSON(response(http.StatusOK, "the crow makes tokens", map[string]any{"token": token}))
	}
}

// Register
func (us *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user model.User
		err := c.Bind(&user)
		if err != nil {
			return c.JSON(response(http.StatusBadRequest, "An error occurred while reading the input"))
		}

		if user.Name == "" || user.Email == "" || user.Password == "" {
			return c.JSON(response(http.StatusBadRequest, "Name, Email, and Password cannot be empty"))
		}

		isNotConflict, err := us.Model.AddUser(user)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(response(http.StatusInternalServerError, "An error occurred while checking the user"))
		} else if !isNotConflict {
			return c.JSON(response(http.StatusConflict, "Someone has already used this email"))
		} else {
			return c.JSON(response(http.StatusOK, "success create user"))
		}
	}
}
