package api

import (
	"fmt"
	"net/http"
	"time"
	"uvent/database"
	"uvent/models"

	"github.com/labstack/echo/v4"
)

type CreateEventForm struct {
	Category string `json:"category" validate:"required"`
	Title  string `json:"title" validate:"required"`
	Detail  string `json:"detail" validate:"required"`
	StartTime  string `json:"start_time" validate:"required"`
	EndTime  string `json:"end_time" validate:"required"`
	Place string `json:"place" validate:"required"`
	Address  string `json:"address" validate:"required"`
	ApplicationDeadline string `json:"application_deadline" validate:"required"`
	ImageURL string `json:"image_url"`
}

func CreateEvent(c echo.Context) error {
	cookie, err := c.Cookie("email")
	if err != nil {
		c.JSON(400, "bad request")
	}
	email := cookie.Value
	var user models.User
	res := database.DB.Where("email = ?", email).First(&user)
	if res.Error != nil {
		return c.JSON(500, "internal server error")
	}
	form := new(CreateEventForm)
	err = c.Bind(&form)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	startTime, err := time.Parse(time.RFC3339, form.StartTime)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "birth must be in format '2006-01-02T15:04:05Z'")
	}
	endTime, err := time.Parse(time.RFC3339, form.EndTime)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "birth must be in format '2006-01-02T15:04:05Z'")
	}
	if !endTime.After(startTime) {
		return c.JSON(http.StatusBadRequest, "end_time is earlier than start_time")
	}
	applicationDeadline, err := time.Parse(time.RFC3339, form.ApplicationDeadline)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "birth must be in format '2006-01-02T15:04:05Z'")
	}
	event := models.Event{
		Category: form.Category,
		UserID: user.ID,
		Title: form.Title,
		Detail: form.Detail,
		StartTime: startTime,
		EndTime: endTime,
		Place: form.Place,
		Address: form.Address,
		ApplicationDeadline: applicationDeadline,
		ImageURL: form.ImageURL,
	}
	database.DB.Create(&event)

	return c.JSON(200, event)
}

func GetEvent(c echo.Context) error {
	return nil
}

func ApplyToEvent(c echo.Context) error {
	return nil
}