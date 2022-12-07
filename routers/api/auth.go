package api

import (
	"net/http"
	"time"
	"uvent/database"
	"uvent/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignupForm struct {
	Name                 string `json:"name" validate:"required"`
	Birth                string `json:"birth"`
	Address              string `json:"address"`
	Gender               string `json:"gender" validate:"required"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,min=8"`
}

func Signup(c echo.Context) error {
	form := new(SignupForm)
	err := c.Bind(&form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if form.Password != form.PasswordConfirmation {
		return c.JSON(http.StatusBadRequest, "password and password_confirmation must be same")
	}
	if len(form.Password) < 8 {
		return c.JSON(http.StatusBadRequest, "password must be at least 8 characters")
	}
	if form.Gender != "male" && form.Gender != "female" && form.Gender != "other" {
		return c.JSON(http.StatusBadRequest, "gender must be 'male' or 'female' or 'other'")
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(form.Password), 8)
	dateFormat := "2006-01-02"
	birth, err := time.Parse(dateFormat, form.Birth)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "birth must be in format '2006-01-02'")
	}
	user := models.User{
		Email:    form.Email,
		Password: password,
		Name:     form.Name,
		Birth:    birth,
		Address:  form.Address,
		Gender:   form.Gender,
	}
	database.DB.Create(&user)

	return c.JSON(200, user)
}

func Login(c echo.Context) error {
	// do something
	return c.JSON(200, "login")
}
