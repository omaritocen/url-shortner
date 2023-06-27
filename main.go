package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"url-shortner/compression"
	"url-shortner/db"
)

func main() {

	db.Init()

	r := gin.Default()

	r.POST("/shorten", compression.ShortenUrl)
	r.GET("/tiny/:shortUrl", compression.GetUrl)

	err := r.Run()
	if err != nil {
		os.Exit(-1)
	}
}
