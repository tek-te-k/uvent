package api

import "github.com/labstack/echo/v4"

func Signup(c echo.Context) error {
	// do something
	return c.JSON(200, "signup")
}

func Login(c echo.Context) error {
	// do something
	return c.JSON(200, "login")
}
