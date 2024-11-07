package aplikasi

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AplikasiController interface {
	FindAll(ctx *gin.Context)
	FindByKd(ctx *gin.Context)
	FindByLimit(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controllerAplikasi struct {
	service AplikasiService
}

// NewAplikasiController creates a new instance of AplikasiController
func NewAplikasiController(service AplikasiService) AplikasiController {
	return &controllerAplikasi{
		service: service,
	}
}

// FindAll returns all aplikasi records or an error if one occurs
func (c *controllerAplikasi) FindAll(ctx *gin.Context) {
	aplikasis, err := c.service.FindAll()
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve aplikasi data")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": aplikasis})
}

// FindByKd returns an aplikasi by kd
func (c *controllerAplikasi) FindByKd(ctx *gin.Context) {
	kdParam := ctx.Param("kd")
	kd, err := strconv.ParseInt(kdParam, 10, 16)
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, "Invalid kd parameter")
		return
	}

	aplikasi, err := c.service.FindByKd(int16(kd))
	if err != nil {
		c.respondWithError(ctx, http.StatusNotFound, "Aplikasi not found")
		return
	}

	// Exclude CreatedAt and UpdatedAt from the response
	aplikasi.CreatedAt = time.Time{}
	aplikasi.UpdatedAt = time.Time{}
	ctx.JSON(http.StatusOK, gin.H{"data": aplikasi})
}

// FindByLimit returns a limited number of aplikasi records based on the limit parameter
func (c *controllerAplikasi) FindByLimit(ctx *gin.Context) {
	limitParam := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit <= 0 {
		c.respondWithError(ctx, http.StatusBadRequest, "Limit must be a positive integer")
		return
	}

	aplikasis, err := c.service.FindByLimit(limit)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve aplikasi data")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": aplikasis})
}

// Create creates a new aplikasi
func (c *controllerAplikasi) Create(ctx *gin.Context) {
	var aplikasi Aplikasi
	if err := ctx.ShouldBindJSON(&aplikasi); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	// Validate required fields
	if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
		c.respondWithError(ctx, http.StatusBadRequest, "All fields are required")
		return
	}

	// Set CreatedAt and UpdatedAt to the current time
	aplikasi.CreatedAt = time.Now()
	aplikasi.UpdatedAt = time.Now()

	result, err := c.service.Create(aplikasi)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, "Failed to create aplikasi")
		return
	}
	// Exclude CreatedAt and UpdatedAt from the response
	result.CreatedAt = time.Time{}
	result.UpdatedAt = time.Time{}
	ctx.JSON(http.StatusCreated, gin.H{"data": result})
}

// Update updates an existing aplikasi
func (c *controllerAplikasi) Update(ctx *gin.Context) {
	var aplikasi Aplikasi
	if err := ctx.ShouldBindJSON(&aplikasi); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	kdParam := ctx.Param("kd")
	kd, err := strconv.ParseInt(kdParam, 10, 16)
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, "Invalid kd parameter")
		return
	}

	// Set kd from path parameter
	aplikasi.Kd = int16(kd)

	// Find existing aplikasi by kd
	_, err = c.service.FindByKd(aplikasi.Kd)
	if err != nil {
		c.respondWithError(ctx, http.StatusNotFound, "Aplikasi not found")
		return
	}

	// Set UpdatedAt to the current time
	aplikasi.UpdatedAt = time.Now()

	err = c.service.Update(aplikasi)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, "Failed to update aplikasi")
		return
	}

	// Exclude CreatedAt and UpdatedAt from the response
	aplikasi.CreatedAt = time.Time{}
	aplikasi.UpdatedAt = time.Time{}
	ctx.JSON(http.StatusOK, gin.H{"message": "Aplikasi updated successfully", "data": aplikasi})
}

// Delete deletes an aplikasi by kd
func (c *controllerAplikasi) Delete(ctx *gin.Context) {
	kdParam := ctx.Param("kd")
	kd, err := strconv.ParseInt(kdParam, 10, 16)
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, "Invalid kd parameter")
		return
	}

	// Find existing aplikasi by kd
	_, err = c.service.FindByKd(int16(kd))
	if err != nil {
		c.respondWithError(ctx, http.StatusNotFound, "Aplikasi not found")
		return
	}

	// Pass only the kd to the Delete method, not the Aplikasi struct
	err = c.service.Delete(int16(kd))
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, "Failed to delete aplikasi")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Aplikasi deleted successfully"})
}

// respondWithError sends a consistent error response
func (c *controllerAplikasi) respondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{"error": message})
}
