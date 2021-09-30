package repositories

import (
	"api-golang/auth"
	"api-golang/models"
	"api-golang/requests"
	"errors"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c echo.Context, userRequest requests.SignIn) error {
	user := new(models.User)
	user.Email = userRequest.Email
	err := user.GetUserByEmail()
	if err != nil {
		return errors.New("NOT_FOUND")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		return errors.New("UNAUTHORIZED")
	}
	if err = auth.GenerateTokensAndSetCookies(user, c); err != nil {
		return errors.New("SERVER_ERROR")
	}

	return nil
}

func SignUp(userRequest requests.SignUp) (*models.User, error) {
	user := new(models.User)
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.FirstName = userRequest.FirstName
	user.LastName = userRequest.LastName
	user.CreateUser()
	return user, nil
}

func GetUser(userID uint) (interface{}, error) {
	user := new(models.User)
	user.ID = userID
	err := user.GetUserByID()
	if err != nil {
		return nil, errors.New("USER_NOT_FOUND")
	}
	return user, nil
}
