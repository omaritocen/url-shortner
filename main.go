package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenUrl)
	r.GET("/tiny/:shortUrl", getUrl)

	err := r.Run()
	if err != nil {
		os.Exit(-1)
	}
}
