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
func (bc *BarangController) Update(c echo.Context) error {
	id := c.Param("id")

	var barang models.Barang
	if err := bc.db.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	newBarang := new(models.Barang)
	if err := c.Bind(newBarang); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	if newBarang.Nama != "" {
		barang.Nama = newBarang.Nama
	}
	if newBarang.Deskripsi != "" {
		barang.Deskripsi = newBarang.Deskripsi
	}
	if newBarang.JumlahStok != 0 {
		barang.JumlahStok = newBarang.JumlahStok
	}
	if newBarang.Harga != 0 {
		barang.Harga = newBarang.Harga
	}
	if newBarang.Kategori != "" {
		barang.Kategori = newBarang.Kategori
	}

	if err := bc.db.Save(&barang).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update Item")
	}

	return c.JSON(http.StatusOK, barang)
}

// Delete a item
func (bc *BarangController) Delete(c echo.Context) error {
	id := c.Param("id")

	var barang models.Barang
	if err := bc.db.First(&barang, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	if err := bc.db.Delete(&barang).Error; err != nil {
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
