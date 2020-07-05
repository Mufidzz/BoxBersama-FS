package main

import (
	"./config"
	"./controller"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "JWT"},
		AllowCredentials:       false,
		ExposeHeaders:          nil,
		MaxAge:                 24 * time.Hour,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))

	image := router.Group("/images")
	{
		image.GET("/:name", controller.GetImage)
		image.POST("/", controller.CreateImage)
	}

	home, _ := os.UserHomeDir()
	err := os.Chdir(filepath.Join(home, "fs-storage", "fs.bb.ofcode.site"))
	if err != nil {
		panic(err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	out, err := os.Create(config.STORAGE_IMAGE_PATH)
	if err != nil {
		panic(err.Error())
		return
	}

	fmt.Println("Working Dir: " + path)
	fmt.Println(out)

	router.Run(":2111")
}
