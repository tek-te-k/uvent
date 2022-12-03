package routers

import (
	"uvent/routers/api"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.POST("/signup", api.Signup)
	app.POST("/login", api.Login)
}
