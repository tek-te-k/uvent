package routers

import (
	api "uvent/routers/api"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	apiv1 := app.Group("/api")
	apiv1.POST("/signup", api.Signup)
	apiv1.POST("/login", api.Login)
}
