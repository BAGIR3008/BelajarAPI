package handler

import (
	"BelajarAPI/features/user"
	"BelajarAPI/helper"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) user.UserController {
	return &controller{
		service: s,
	}
}

func (ct *controller) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}
		err = ct.service.Register(input)
		if err != nil {
			if err.Error() == helper.ErrorGeneralServer {
				return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))
			}
			if strings.Contains(err.Error(), "validation") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorInvalidValidate, map[string]any{"description": strings.Split(err.Error(), "\n")}))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, err.Error()))
		}
		return c.JSON(helper.ResponseFormat(http.StatusCreated, "Congratulations, the data has been registered"))
	}
}

func (ct *controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		var processData user.User
		processData.Email = input.Email
		processData.Password = input.Password

		result, token, err := ct.service.Login(processData)
		if err != nil {
			if err.Error() == helper.ErrorUserCredential {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserCredential))
			} else if strings.Contains(err.Error(), "validation") {
				return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorInvalidValidate, map[string]any{"description": err.Error()}))
			}
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))
		}

		var responseData LoginResponse
		responseData.Email = result.Email
		responseData.Name = result.Name
		responseData.Token = token

		return c.JSON(helper.ResponseFormat(http.StatusOK, "login successful", map[string]any{"data": responseData}))

	}
}

func (ct *controller) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		profile, err := ct.service.Profile(token)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))
		}

		var profileResponse ProfileResponse
		if err := helper.Recast(&profile, &profileResponse); err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success get user profile", map[string]any{"user": profileResponse}))
	}
}
