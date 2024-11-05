package main

import (
	"rsudlampung/helper"
	"rsudlampung/modules"

	"log"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	os.Setenv("TZ", "Asia/Jakarta")
	configEnv, errEnv := helper.LoadConfig("./")
	if errEnv != nil {
		log.Fatal("cannot load config:", errEnv)
	}

	logFile := configEnv.LogFile
	if logFile == "on" {
		helper.SetupLogOutput()
	}

	ginModeEnv := configEnv.GinMode
	if ginModeEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	mainServer := gin.New()
	mainServer.Use(gin.Recovery())

	config := cors.DefaultConfig()
	origins := strings.Split(configEnv.AllowOrigin, ",")
	if configEnv.Env == "staging" {
		origins = append(origins, "http://localhost:3000")
	}
	config.AllowOrigins = origins

	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "DELETE"}
	config.AllowHeaders = []string{"Access-Control-Allow-Origin", "authorization", "content-type", "content-length", "user-agent", "Host", "accept"}
	log.Println("config", config.AllowOrigins)
	mainServer.Use(cors.New(config))
	mainServer.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "root path")
	})

	m := modules.NewVersion(configEnv, mainServer)
	m.Run()

	port := configEnv.Port
	mainServer.Run(":" + port)

}
