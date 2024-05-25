package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type User struct {
	gorm.Model
	Name string
}

func InitUserRoutes(e *echo.Echo, db *gorm.DB) {
	e.GET("/users", func(c echo.Context) error {
		return GetUsers(c, db)
	})
}

func GetUsers(c echo.Context, db *gorm.DB) error {
	return c.String(http.StatusOK, " GET  Users route")
}
