package controller

import (
	"../config"
	"../structs"
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func GetImage(c *gin.Context) {
	fileId := c.Param("name")
	fmt.Println(config.STORAGE_IMAGE_PATH + fileId)
	c.File(config.STORAGE_IMAGE_PATH + fileId + ".png")
}

func CreateImage(c *gin.Context) {
	var (
		images []structs.Image
	)

	err := c.BindJSON(&images)

	if err != nil {
		fmt.Println("ERR JSON BIND")
	}

	for _, img := range images {
		base64pure := img.Base
		imageName := img.Name

		coi := strings.Index(base64pure, ",")
		base64img := base64pure[coi+1:]

		decoded, _ := base64.StdEncoding.DecodeString(base64img)
		img, _, _ := image.Decode(bytes.NewReader(decoded))

		out, err := os.Create(config.STORAGE_IMAGE_PATH + imageName + ".png")

		if err != nil {
			//TODO: Handle Error Here
		}

		switch strings.TrimSuffix(base64pure[5:coi], ";base64") {
		case "image/png":
			err = png.Encode(out, img)
			if err != nil {
				//TODO: Handle Error Here
			}
			break
		case "image/jpeg":
			err = jpeg.Encode(out, img, &jpeg.Options{Quality: 80})
			if err != nil {
				//TODO: Handle Error Here
			}
			break
		}
	}

	c.JSON(200, gin.H{
		"message": "Success",
	})
}
