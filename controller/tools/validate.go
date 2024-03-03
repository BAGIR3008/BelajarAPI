package tools

import (
	"fmt"

	"github.com/go-playground/validator"
)

var validate = validator.New()

func Validate(s interface{}) []string {
	var messages = []string{}
	err := validate.Struct(s)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// fmt.Println(e.ActualTag(), e.Field(), e.Kind(), e.Namespace(), e.Param(), e.StructField(), e.StructNamespace(), e.Tag(), e.Value(), e.Kind().String(), e.Type().Align())
			if e.Param() != "" {
				messages = append(messages, fmt.Sprintf("'%s' => %s must be %s", e.Value(), e.ActualTag(), e.Param()))
			} else if e.ActualTag() == "required" {
				messages = append(messages, fmt.Sprintf("%s => is required", e.Field()))
			} else {
				messages = append(messages, fmt.Sprintf("'%s' => is not %s", e.Value(), e.ActualTag()))
			}
		}
	}
	return messages
}
