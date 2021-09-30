package userValidators

import (
	"api-golang/requests"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type (
	SignUpValidator struct {
		validator *validator.Validate
	}
)

func (v *SignUpValidator) New() {
	v.validator = validator.New()
	v.validator.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6 && len(fl.Field().String()) < 16
	})
}

func (v *SignUpValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		_, ok := err.(*echo.HTTPError)
		errorString := ""
		if !ok {
			errorString = err.Error()
		}
		fieldsError := make(map[string]string)
		signUpRequest := i.(*requests.SignUp)
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for index, err := range castedObject {
				message := ""
				log.Print(err.Field(), err)
				switch err.Tag() {
				case "required":
					message = fmt.Sprintf("%s_IS_REQUIRED",
						strings.ToUpper(err.Field()))
				case "email":
					message = fmt.Sprintf("%s_IS_NOT_VALID_EMAIL",
						strings.ToUpper(err.Field()))
				case "passwd":
					message = fmt.Sprintf("%s_INVALID_FORMAT",
						strings.ToUpper(err.Field()))
				}
				fieldsError[signUpRequest.GetJSONField(index)] = message
			}
			fieldsErrorString, _ := json.Marshal(fieldsError)

			errorString = string(fieldsErrorString)
		}
		return errors.New(errorString)
	}
	return nil
}
