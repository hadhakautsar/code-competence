package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"code-competence-remidi/models"
)

type BarangController struct {
	db *gorm.DB
}

// Constructor for BarangController
func NewBarangController(db *gorm.DB) *BarangController {
	return &BarangController{db}
}

// Create a barang
func (bc *BarangController) Create(c echo.Context) error {
	barang := new(models.Barang)
	if err := c.Bind(barang); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if err := bc.db.Create(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create item")
	}

	return c.JSON(http.StatusCreated, barang)
}

// Read all barang
func (bc *BarangController) ReadAll(c echo.Context) error {
	var barangs []models.Barang
	if err := bc.db.Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get item")
	}

	return c.JSON(http.StatusOK, barangs)
}

// Read a barang
func (bc *BarangController) Read(c echo.Context) error {
	id := c.Param("id")

	var barang models.Barang
	if err := bc.db.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	return c.JSON(http.StatusOK, barang)
}

// Update a item
/* func (bc *BarangController) Update(c echo.Context) error {
	id := c.Param("id")

	var member models.Barang
	if err := bc.db.First(&member, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	newMember := new(models.Barang)
	if err := c.Bind(newMember); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if newMember.Username != "" {
		member.Username = newMember.Username
	}
	if newMember.Email != "" {
		member.Email = newMember.Email
	}
	if newMember.Password != "" {
		member.Password = newMember.Password
	}
	if newMember.Handicap != 0 {
		member.Handicap = newMember.Handicap
	}
	if newMember.Score != 0 {
		member.Score = newMember.Score
	}

	if err := bc.db.Save(&member).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update Item")
	}

	return c.JSON(http.StatusOK, member)
}
*/
// Delete a item
func (bc *BarangController) Delete(c echo.Context) error {
	id := c.Param("id")

	var member models.Barang
	if err := bc.db.First(&member, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	if err := bc.db.Delete(&member).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete Item")
	}

	return c.JSON(http.StatusOK, "Item deleted successfully")
}

// Read all item by category
func (bc *BarangController) ReadByCategory(c echo.Context) error {
	kategoriID := c.Param("kategoriID")

	var barangs []models.Barang
	if err := bc.db.Where("kategori = ?", kategoriID).Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get item")
	}

	return c.JSON(http.StatusOK, barangs)
}

// Read barang by keyword
func (bc *BarangController) ReadByKeyword(c echo.Context) error {
	keyword := c.QueryParam("keyword")

	var barangs []models.Barang
	if err := bc.db.Where("nama LIKE ?", "%"+keyword+"%").Find(&barangs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get barang")
	}

	return c.JSON(http.StatusOK, barangs)
}
