package routers

import (
	"github.com/labstack/echo/v4"
)

func Moke(c echo.Context) error {
	return c.JSON(200, "mokemoke")
}
