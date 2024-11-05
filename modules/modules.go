package modules

import (
	"rsudlampung/helper"

	// aplikasi "rsudlampung/modules/akses_pengguna/aplikasi"
	// fitur "rsudlampung/modules/akses_pengguna/fitur"
	// group "rsudlampung/modules/akses_pengguna/group"
	// modul "rsudlampung/modules/akses_pengguna/modul"
	// pengguna "rsudlampung/modules/akses_pengguna/pengguna"
	// group_aplikasi "rsudlampung/modules/akses_pengguna/group_aplikasi"
	// group_fitur "rsudlampung/modules/akses_pengguna/group_fitur"
	// group_modul "rsudlampung/modules/akses_pengguna/group_modul"

	"github.com/gin-gonic/gin"
)

type Versions interface {
	Run()
}

type versions struct {
	configEnv  helper.Config
	mainServer *gin.Engine
}

func NewVersion(configEnv helper.Config, mainServer *gin.Engine) Versions {
	return &versions{
		configEnv:  configEnv,
		mainServer: mainServer,
	}
}

func (s *versions) Run() {
	// apiSistemRoutes := s.mainServer.Group("/sistem")
	// db_aksesPengguna := helper.OpenDB(s.configEnv.DB, s.configEnv.SCHEMA, "v010")

	// aplikasi := aplikasi.NewAplikasiServer(apiSistemRoutes, db_aksesPengguna, "v010")
	// aplikasi.Init()

}
