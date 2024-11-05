package aplikasi

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AplikasiController interface {
	FindAll() []Aplikasi
	FindByKd(ctx *gin.Context) (Aplikasi, error)
	Create(ctx *gin.Context) (Aplikasi, error)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	FindByLimit(ctx *gin.Context) []Aplikasi
}

type controllerAplikasi struct {
	service AplikasiService
}

func NewAplikasiController(db *gorm.DB) AplikasiController {
	return &controllerAplikasi{
		service: NewAplikasiService(db),
	}
}

func (c *controllerAplikasi) FindAll() []Aplikasi {
	return c.service.FindAll()
}

func (c *controllerAplikasi) FindByKd(ctx *gin.Context) (Aplikasi, error) {
	var aplikasi Aplikasi
	kdStr := ctx.Param("kd")

	kd, err := strconv.Atoi(kdStr)
	if err != nil {
		return Aplikasi{}, errors.New("kd must be an integer")
	}

	aplikasi = c.service.FindByKd(kd)
	if (aplikasi == Aplikasi{}) {
		return Aplikasi{}, errors.New("Aplikasi tidak valid")
	}
	return aplikasi, nil
}

func (c *controllerAplikasi) Create(ctx *gin.Context) (Aplikasi, error) {
	var aplikasi Aplikasi
	err := ctx.ShouldBindJSON(&aplikasi)
	if err != nil {
		return Aplikasi{}, err
	}
	result, err := c.service.Create(aplikasi)
	if err != nil {
		return Aplikasi{}, err
	}
	return result, nil
}

func (c *controllerAplikasi) Update(ctx *gin.Context) error {
	var aplikasi Aplikasi
	err := ctx.ShouldBindJSON(&aplikasi)
	if err != nil {
		return err
	}

	err = c.service.Update(aplikasi)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerAplikasi) Delete(ctx *gin.Context) error {
	kdStr := ctx.Param("kd")

	kd, err := strconv.Atoi(kdStr)
	if err != nil {
		return errors.New("kd must be an integer")
	}

	aplikasi := c.service.FindByKd(kd)
	if (aplikasi == Aplikasi{}) {
		return errors.New("Aplikasi tidak valid")
	}

	aplikasi.Kd = kd
	err = c.service.Delete(aplikasi)
	if err != nil {
		return err
	}
	return nil
}

func (c *controllerAplikasi) FindByLimit(ctx *gin.Context) []Aplikasi {
	offsetStr := ctx.Query("offset")
	limitStr := ctx.Query("limit")

	// Convert offset and limit from string to int
	offsetInt := 0 // Default value
	limitInt := 10 // Default value
	var err error

	if offsetStr != "" {
		offsetInt, err = strconv.Atoi(offsetStr)
		if err != nil {
			offsetInt = 0 // Set to default if conversion fails
		}
	}

	if limitStr != "" {
		limitInt, err = strconv.Atoi(limitStr)
		if err != nil {
			limitInt = 10 // Set to default if conversion fails
		}
	}

	return c.service.FindByLimit(offsetInt, limitInt)
}
