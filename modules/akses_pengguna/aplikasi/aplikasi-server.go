package aplikasi

import (
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
	// Create AplikasiService with the database
	aplikasiService := NewAplikasiService(s.database)

	// Initialize the controller with the AplikasiService
	aplikasiControl := NewAplikasiController(aplikasiService)

	// Endpoint for fetching all aplikasi records
	s.apiRoutes.GET("/"+s.version+"/aplikasi", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		aplikasiControl.FindAll(ctx)
	})

	// Endpoint for fetching an aplikasi by kd
	s.apiRoutes.GET("/"+s.version+"/aplikasi/:kd", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		aplikasiControl.FindByKd(ctx)
	})

	// Endpoint for fetching aplikasi with a limit
	s.apiRoutes.GET("/"+s.version+"/aplikasi/limit", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		aplikasiControl.FindByLimit(ctx)
	})

	// Endpoint for creating a new aplikasi
	s.apiRoutes.POST("/"+s.version+"/aplikasi", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		aplikasiControl.Create(ctx)
	})

	// Endpoint for updating an aplikasi by kd
	s.apiRoutes.PUT("/"+s.version+"/aplikasi/:kd", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		aplikasiControl.Update(ctx)
	})

	// Endpoint for deleting an aplikasi by kd
	s.apiRoutes.DELETE("/"+s.version+"/aplikasi/:kd", mid_auth.BasicAuth(), func(ctx *gin.Context) {
		aplikasiControl.Delete(ctx)
	})
}
