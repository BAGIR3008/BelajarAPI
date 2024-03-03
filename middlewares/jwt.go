package middlewares

import (
	"BelajarAPI/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string) (string, error) {
	var data = jwt.MapClaims{}
	data["email"] = email
	data["iat"] = time.Now()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var proccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := proccessToken.SignedString([]byte(config.JWTSECRET))
	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}

func DecodeToken(i interface{}) string {
	var claim = i.(*jwt.Token).Claims.(jwt.MapClaims)
	var result string

	if val, found := claim["email"]; found {
		result = val.(string)
	}

	return result
}
