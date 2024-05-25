package router


import (
	"net/http"
	"github.com/labstack/echo/v4"
		"gorm.io/gorm"
)




func InitRoutes(e *echo.Echo,db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	InitUserRoutes(e,db)
}
