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
	aplikasiControl := NewAplikasiController(NewAplikasiService(s.database))

	s.apiRoutes.GET("/"+s.version+"/aplikasi", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, aplikasiControl.FindAll(ctx))
	})

	s.apiRoutes.GET("/"+s.version+"/aplikasi/:kd", func(ctx *gin.Context) {
		result, err := aplikasiControl.FindByKd(ctx)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result})
		}
	})

	s.apiRoutes.POST("/"+s.version+"/aplikasi", func(ctx *gin.Context) {
		result, err := aplikasiControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"data": result})
		}
	})

	s.apiRoutes.PUT("/"+s.version+"/aplikasi", func(ctx *gin.Context) {
		err := aplikasiControl.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Aplikasi updated successfully"})
		}
	})

	s.apiRoutes.DELETE("/"+s.version+"/aplikasi/:kd", func(ctx *gin.Context) {
		err := aplikasiControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Aplikasi deleted successfully"})
		}
	})
}
