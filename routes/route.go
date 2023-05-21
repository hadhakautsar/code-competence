package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"code-competence-remidi/controllers"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	BarangController := controllers.NewBarangController(db)

	// Admin routes
	e.POST("/items", BarangController.Create)
	e.GET("/items", BarangController.ReadAll)
	e.GET("/items/:id", BarangController.Read)
	/* e.PUT("/items/:id", BarangController.Update) */
	e.DELETE("/items/:id", BarangController.Delete)
	e.GET("/items/category/:kategoriID", BarangController.ReadByCategory)
	e.GET("/items", BarangController.ReadByKeyword)
}
