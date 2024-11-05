package aplikasi

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AplikasiController interface {
	FindAll() []Aplikasi
	FindByKd(ctx *gin.Context) (Aplikasi, error)
	FindByLimit(ctx *gin.Context) ([]Aplikasi, error)
	Create(ctx *gin.Context) (Aplikasi, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controllerAplikasi struct {
	service AplikasiService
}

// NewAplikasiController creates a new instance of AplikasiController
func NewAplikasiController(db *gorm.DB) AplikasiController {
	return &controllerAplikasi{
		service: NewAplikasiService(db),
	}
}

// FindAll returns all aplikasi records
func (c *controllerAplikasi) FindAll() []Aplikasi {
	return c.service.FindAll()
}

// FindByKd returns an aplikasi by kd
func (c *controllerAplikasi) FindByKd(ctx *gin.Context) (Aplikasi, error) {
	kd := ctx.Param("kd")
	aplikasi, err := c.service.FindByKd(kd)
	if err != nil {
		return Aplikasi{}, errors.New("Aplikasi tidak ditemukan")
	}
	return aplikasi, nil
}

// FindByLimit returns a limited number of aplikasi records based on the limit parameter
func (c *controllerAplikasi) FindByLimit(ctx *gin.Context) ([]Aplikasi, error) {
	limitParam := ctx.Query("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		return nil, errors.New("limit harus berupa angka positif")
	}

	aplikasis, err := c.service.FindByLimit(limit)
	if err != nil {
		return nil, err
	}
	return aplikasis, nil
}

// Create creates a new aplikasi
func (c *controllerAplikasi) Create(ctx *gin.Context) (Aplikasi, error) {
	var aplikasi Aplikasi
	if err := ctx.ShouldBindJSON(&aplikasi); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data."}) // Return if binding fails
		return Aplikasi{}, err
	}

	// Validate the required fields
	if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi."}) // All fields must be filled
		return Aplikasi{}, errors.New("semua field harus diisi")
	}

	result, err := c.service.Create(aplikasi)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create aplikasi."}) // Handle errors during creation
		return Aplikasi{}, err
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": result}) // Return the created aplikasi
	return result, nil
}

// Update updates an existing aplikasi
func (c *controllerAplikasi) Update(ctx *gin.Context) error {
    var aplikasi Aplikasi

    // Bind the incoming JSON to the Aplikasi struct
    if err := ctx.ShouldBindJSON(&aplikasi); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data."})
        return err
    }

    // Validate the required fields
    if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Semua field harus diisi."})
        return errors.New("semua field harus diisi")
    }

    // Call the service to update the aplikasi
    err := c.service.Update(aplikasi)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update aplikasi."})
        return err
    }

    // Ensure that you return a success message
    ctx.JSON(http.StatusOK, gin.H{"message": "Aplikasi updated successfully."})
    return nil
}

// Delete deletes an aplikasi by kd
func (c *controllerAplikasi) Delete(ctx *gin.Context) error {
	kd := ctx.Param("kd")

	aplikasi, err := c.service.FindByKd(kd)
	if err != nil || aplikasi.Kd == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Aplikasi tidak ditemukan."}) // Return error if not found
		return errors.New("Aplikasi tidak ditemukan") // Also return an error
	}

	err = c.service.Delete(aplikasi)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete aplikasi."}) // Handle errors during deletion
		return err // Return the error
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Aplikasi deleted successfully."}) // Return success message
	return nil // Return nil if successful
}
