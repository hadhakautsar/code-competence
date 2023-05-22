package main

import (
	config "code-competence-remidi/configs"
	"code-competence-remidi/middleware"
	"code-competence-remidi/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Initialize GORM database
	db, err := config.InitDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Use(middleware.LogMiddleware)

	// Register routes
	routes.InitRoutes(e, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start("0.0.0.0:" + port))
}
