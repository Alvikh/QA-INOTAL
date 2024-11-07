package aplikasi

import (
	"math"
	"net/http"
	"strconv"
	"strings"
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

// Constants for error messages
const (
	ErrInvalidInput     = "Invalid input data"
	ErrAplikasiNotFound = "Aplikasi not found"
	ErrFailedToCreate   = "Failed to create aplikasi"
	ErrInvalidKd        = "Invalid kd parameter"
	ErrInvalidLimit     = "Invalid limit parameter"
	ErrInvalidOffset    = "Invalid offset parameter"
	ErrInvalidRange     = "Invalid kd range"
	ErrSQLInjection     = "Invalid input data (SQL Injection detected)"
)

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
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidKd)
		return
	}

	// Validate kd range
	if kd < math.MinInt16 || kd > math.MaxInt16 {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidRange)
		return
	}

	aplikasi, err := c.service.FindByKd(int16(kd))
	if err != nil {
		c.respondWithError(ctx, http.StatusNotFound, ErrAplikasiNotFound)
		return
	}

	// Exclude CreatedAt and UpdatedAt from the response
	aplikasi.CreatedAt = time.Time{}
	aplikasi.UpdatedAt = time.Time{}
	ctx.JSON(http.StatusOK, gin.H{"data": aplikasi})
}

// FindByLimit returns a limited number of aplikasi records based on the limit and offset parameters
func (c *controllerAplikasi) FindByLimit(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidLimit)
		return
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidOffset)
		return
	}

	aplikasis, err := c.service.FindByLimit(limit, offset)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve aplikasi data")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": aplikasis})
}

// sanitizeInput sanitizes the input string to prevent SQL injection
func sanitizeInput(input string) string {
	badWords := []string{"DROP", "DELETE", "INSERT", "UPDATE", "TRUNCATE", "ALTER", "CREATE", "--", ";", "/*", "*/"}
	for _, word := range badWords {
		if strings.Contains(strings.ToUpper(input), word) {
			return ""
		}
	}
	replacer := strings.NewReplacer("'", "", "\"", "", ";", "", "--", "")
	return replacer.Replace(input)
}

// Create a new aplikasi record
func (c *controllerAplikasi) Create(ctx *gin.Context) {
	var aplikasi Aplikasi
	if err := ctx.ShouldBindJSON(&aplikasi); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidInput)
		return
	}

	if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
		c.respondWithError(ctx, http.StatusBadRequest, "All fields are required")
		return
	}

	// Sanitize input to prevent SQL injection
	aplikasi.Nama = sanitizeInput(aplikasi.Nama)
	aplikasi.Label = sanitizeInput(aplikasi.Label)
	aplikasi.Logo = sanitizeInput(aplikasi.Logo)
	aplikasi.UrlFE = sanitizeInput(aplikasi.UrlFE)
	aplikasi.UrlAPI = sanitizeInput(aplikasi.UrlAPI)

	if aplikasi.Nama == "" || aplikasi.Label == "" || aplikasi.Logo == "" || aplikasi.UrlFE == "" || aplikasi.UrlAPI == "" {
		c.respondWithError(ctx, http.StatusBadRequest, ErrSQLInjection)
		return
	}

	aplikasi.CreatedAt = time.Now()
	aplikasi.UpdatedAt = time.Now()

	result, err := c.service.Create(aplikasi)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, ErrFailedToCreate)
		return
	}

	// Exclude CreatedAt and UpdatedAt from the response
	result.CreatedAt = time.Time{}
	result.UpdatedAt = time.Time{}
	ctx.JSON(http.StatusCreated, gin.H{"data": result})
}

// Update an existing aplikasi record
func (c *controllerAplikasi) Update(ctx *gin.Context) {
	var aplikasi Aplikasi
	if err := ctx.ShouldBindJSON(&aplikasi); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidInput)
		return
	}

	kdParam := ctx.Param("kd")
	kd, err := strconv.ParseInt(kdParam, 10, 16)
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidKd)
		return
	}

	// Validate kd range
	if kd < math.MinInt16 || kd > math.MaxInt16 {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidRange)
		return
	}

	aplikasi.Kd = int16(kd)

	_, err = c.service.FindByKd(aplikasi.Kd)
	if err != nil {
		c.respondWithError(ctx, http.StatusNotFound, ErrAplikasiNotFound)
		return
	}

	// Sanitize input to prevent SQL injection
	aplikasi.Nama = sanitizeInput(aplikasi.Nama)
	aplikasi.Label = sanitizeInput(aplikasi.Label)
	aplikasi.Logo = sanitizeInput(aplikasi.Logo)
	aplikasi.UrlFE = sanitizeInput(aplikasi.UrlFE)
	aplikasi.UrlAPI = sanitizeInput(aplikasi.UrlAPI)

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

// Delete an aplikasi record by kd
func (c *controllerAplikasi) Delete(ctx *gin.Context) {
	kdParam := ctx.Param("kd")
	kd, err := strconv.ParseInt(kdParam, 10, 16)
	if err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidKd)
		return
	}

	// Validate kd range
	if kd < math.MinInt16 || kd > math.MaxInt16 {
		c.respondWithError(ctx, http.StatusBadRequest, ErrInvalidRange)
		return
	}

	_, err = c.service.FindByKd(int16(kd))
	if err != nil {
		c.respondWithError(ctx, http.StatusNotFound, ErrAplikasiNotFound)
		return
	}

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
