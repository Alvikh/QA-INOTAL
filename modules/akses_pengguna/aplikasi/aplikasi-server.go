package aplikasi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AplikasiServer interface {
	Init()
}

type aplikasiServer struct {
	apiRoutes *gin.RouterGroup
	database  *gorm.DB
	version   string
}

func NewAplikasiServer(apiR *gin.RouterGroup, db *gorm.DB, ver string) AplikasiServer {
	return &aplikasiServer{
		apiRoutes: apiR,
		database:  db,
		version:   ver,
	}
}

func (s *aplikasiServer) Init() {
	aplikasiControl := NewAplikasiController(s.database)

	// Route to get all applications
	s.apiRoutes.GET("/"+s.version+"/aplikasi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, aplikasiControl.FindAll())
	})

	// Route to get an application by kd
	s.apiRoutes.GET("/"+s.version+"/aplikasi/:kd", func(ctx *gin.Context) {
		result, err := aplikasiControl.FindByKd(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	// Route to create a new application
	s.apiRoutes.POST("/"+s.version+"/aplikasi", func(ctx *gin.Context) {
		result, err := aplikasiControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	// Route to update an existing application
	s.apiRoutes.PUT("/"+s.version+"/aplikasi", func(ctx *gin.Context) {
		err := aplikasiControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})

	// Route to delete an application by kd
	s.apiRoutes.DELETE("/"+s.version+"/aplikasi/:kd", func(ctx *gin.Context) {
		err := aplikasiControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})

	// Route to find applications with pagination
	s.apiRoutes.GET("/"+s.version+"/aplikasi/limit", func(ctx *gin.Context) {
		data := aplikasiControl.FindByLimit(ctx)
		ctx.JSON(http.StatusOK, gin.H{"data": data, "error": nil})
	})
}
