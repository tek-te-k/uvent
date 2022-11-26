package main

import (
	// "uvent/database"
	"uvent/routers"

	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	// TODO DB用のdocker-composeを作成する
	// database.Connect()
	app := echo.New()
	// app.Validator = &validate.Validator{Validator: validator.New()}
	// routers.SetupRouter(app)
	routers.SetupRoutes(app)
	app.Logger.Fatal(app.Start(":8080"))
}
