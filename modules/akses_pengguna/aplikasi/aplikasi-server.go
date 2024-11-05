package aplikasi

import (
	"net/http"

	"rsudlampung/middlewares/mid_auth"

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

	// Endpoint for fetching all aplikasi records
	s.apiRoutes.GET("/"+s.version+"/aplikasi", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, aplikasiControl.FindAll())
	})

	// Endpoint for fetching an aplikasi by ID
	s.apiRoutes.GET("/"+s.version+"/aplikasi/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := aplikasiControl.FindByKd(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"data": result, "error": nil})
		}
	})

	// Endpoint for creating a new aplikasi
	s.apiRoutes.POST("/"+s.version+"/aplikasi", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		result, err := aplikasiControl.Create(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()}) // Return 400 for validation errors
		} else {
			ctx.JSON(http.StatusCreated, gin.H{"data": result, "error": nil}) // Return 201 for successful creation
		}
	})

// Endpoint for updating an aplikasi
s.apiRoutes.PUT("/"+s.version+"/aplikasi/:kd", mid_auth.BasicAuth(), func(ctx *gin.Context) {
	err := aplikasiControl.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Return 400 for validation errors
	} else {
		ctx.JSON(http.StatusOK, gin.H{"error": nil}) // Return 200 for successful updates
	}
})


	// Endpoint for deleting an aplikasi
	s.apiRoutes.DELETE("/"+s.version+"/aplikasi/:id", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		err := aplikasiControl.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"error": nil})
		}
	})
}
