package main

import (
	"os"
	"url-shortner/model"

	"github.com/gin-gonic/gin"
)

func main() {

	db := model.Database()
	db.DB()

	r := gin.Default()

	r.POST("/shorten", shortenUrl)
	r.GET("/tiny/:shortUrl", getUrl)

	err := r.Run()
	if err != nil {
		os.Exit(-1)
	}
}
