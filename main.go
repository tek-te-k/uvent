package main

import (
	// "uvent/database"
	"uvent/routers"


	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func setupRouter(app *echo.Echo) {
	app.GET("/", routers.Moke)
}

func main() {
	// TODO DB用のdocker-composeを作成する
	// database.Connect()
	app := echo.New()
	// app.Validator = &validate.Validator{Validator: validator.New()}
	// routers.SetupRouter(app)
	setupRouter(app)
	app.Logger.Fatal(app.Start(":8080"))
}
