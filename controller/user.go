package controller

import (
	"BelajarAPI/controller/message"
	"BelajarAPI/controller/tools"
	"BelajarAPI/middlewares"
	"BelajarAPI/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.Models
}

func (us *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataLogin message.LoginRequest
		err := c.Bind(&dataLogin)
		if err != nil {
			return c.JSON(message.Response(http.StatusBadRequest, "An error occurred while reading the input"))
		}

		errs := tools.Validate(dataLogin)
		if len(errs) != 0 {
			return c.JSON(message.Response(http.StatusBadRequest, "Invalid request", map[string]any{"errors": errs}))
		}

		user, err := us.Model.Login(dataLogin.Email, dataLogin.Password)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while checking the user"))
		} else if user.Name == "" {
			return c.JSON(message.Response(http.StatusBadRequest, "Incorrect email or password"))
		}

		token, err := middlewares.GenerateJWT(user.Email)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "Failed to create token"))
		}

		return c.JSON(message.Response(http.StatusOK, "Successfully created a token", map[string]any{"token": token}))
	}
}

func (us *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var dataRegister message.RegisterRequest
		err := c.Bind(&dataRegister)
		if err != nil {
			return c.JSON(message.Response(http.StatusBadRequest, "An error occurred while reading the input"))
		}

		errs := tools.Validate(dataRegister)
		if len(errs) != 0 {
			return c.JSON(message.Response(http.StatusBadRequest, "Invalid request", map[string]any{"errors": errs}))
		}

		user, err := tools.TypeConverter[model.User](&dataRegister)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "Internal Server Error"))
		}

		isNotConflict, err := us.Model.Register(user)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while checking the user"))
		} else if !isNotConflict {
			return c.JSON(message.Response(http.StatusConflict, "Someone has already used this email"))
		} else {
			return c.JSON(message.Response(http.StatusOK, "success create user"))
		}
	}
}

func (us *UserController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var mailFromToken = middlewares.DecodeToken(c.Get("user"))

		user, err := us.Model.Profile(mailFromToken)
		if err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "An error occurred while calling user data"))
		} else if user.Name == "" {
			return c.JSON(message.Response(http.StatusNotFound, "Email not found"))
		}

		if userResponse, err := tools.TypeConverter[[]message.UserResponse](&user); err != nil {
			log.Println(err.Error())
			return c.JSON(message.Response(http.StatusInternalServerError, "Internal Server Error"))
		} else {
			return c.JSON(message.Response(http.StatusOK, "success get user", map[string]any{"user": userResponse}))
		}
	}
}
