package main

import (
	"github.com/labstack/echo/v4"
	"github.com/tigawanna/mock-github-api/db"
	"github.com/tigawanna/mock-github-api/router"
)

func main() {

	db, err := db.ConnectDatabase()
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&router.User{})

	e := echo.New()

	router.InitRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}
