package controllers

import (
	"api-golang/action/repositories"
	"api-golang/models"
	"api-golang/requests"
	"api-golang/responses"
	"encoding/json"
	"net/http"
	"strconv"

	userValidators "api-golang/validators/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func SignIn(c echo.Context) error {
	signInValidator := new(userValidators.SignInValidator)
	signInValidator.New()
	userRequest := new(requests.SignIn)
	if err := c.Bind(userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, responses.Message{StatusCode: 1400, Message: err.Error()})
	}
	if err := signInValidator.Validate(userRequest); err != nil {
		fields := make(map[string]string)
		json.Unmarshal([]byte(err.Error()), &fields)
		return c.JSON(http.StatusBadRequest, responses.Message{StatusCode: 1400, Message: "SOME_FIELD_INVALID", Fields: fields})
	}
	if err := repositories.SignIn(c, *userRequest); err != nil {
		switch err.Error() {
		case "NOT_FOUND":
			return c.JSON(http.StatusNotFound, responses.Message{StatusCode: 1404, Message: "USER_NOT_FOUND"})
		case "UNAUTHORIZED":
			return c.JSON(http.StatusUnauthorized, responses.Message{StatusCode: 1401, Message: "UNAUTHORIZED"})
		case "SERVER_ERROR":
			return c.JSON(http.StatusInternalServerError, responses.Message{StatusCode: 1500, Message: "SERVER_ERROR"})
		}
	}

	return c.JSON(http.StatusOK, responses.Message{StatusCode: 1200, Message: "SUCCESS"})
}

func SignUp(c echo.Context) error {
	signUpValidator := new(userValidators.SignUpValidator)
	signUpValidator.New()
	userRequest := new(requests.SignUp)
	if err := c.Bind(userRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := signUpValidator.Validate(userRequest); err != nil {
		fields := make(map[string]string)
		json.Unmarshal([]byte(err.Error()), &fields)
		return c.JSON(http.StatusBadRequest, responses.Message{StatusCode: 1400, Message: "SOME_FIELD_INVALID", Fields: fields})
	}
	user, err := repositories.SignUp(*userRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.Message{StatusCode: 1500, Message: "SERVER_ERROR"})
	}
	userResponse := new(responses.User)
	userResponse.ID = user.ID
	userResponse.Email = user.Email
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName
	return c.JSON(http.StatusOK, userResponse)
}

func GetUser(c echo.Context) error {
	cookie, _ := c.Cookie("token")
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, responses.Message{StatusCode: 1401, Message: "UNAUTHORIZED"})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	id, err := strconv.ParseUint(claims.Issuer, 10, 64)
	userInterface, err := repositories.GetUser(uint(id))
	if err != nil {
		switch err.Error() {
		case "NOT_FOUND":
			return c.JSON(http.StatusNotFound, responses.Message{StatusCode: 1404, Message: "USER_NOT_FOUND"})
		}
	}
	user := userInterface.(*models.User)
	userResponse := new(responses.User)
	userResponse.ID = user.ID
	userResponse.Email = user.Email
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName
	return c.JSON(http.StatusOK, userResponse)
}
