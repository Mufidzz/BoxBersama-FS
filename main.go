package main

import (
	"./controller"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
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

	path, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Working Dir: " + path)

	router.Run(":2111")
}
