package requests

import (
	"reflect"
)

type (
	SignUp struct {
		Email     string `json:"email" validate:"required"`
		Password  string `json:"password" validate:"required,passwd"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
	}
	SignIn struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required,passwd"`
	}
)

func (b SignIn) GetJSONField(index int) string {
	val := reflect.ValueOf(b)
	return val.Type().Field(index).Tag.Get("json")
}

func (b SignUp) GetJSONField(index int) string {
	val := reflect.ValueOf(b)
	return val.Type().Field(index).Tag.Get("json")
}
