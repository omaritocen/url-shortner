package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/shorten", func(c *gin.Context) {

		type RequestBody struct {
			LongUrl string `json:"longUrl" binding:"required"`
		}
		var body RequestBody

		if err := c.ShouldBindJSON(&body); err != nil {
			// TODO: Handle Meaningful error message
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		shortUrl := getShortURLFromLongURL(body.LongUrl)

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Tiny url is created",
			"longUrl":  body.LongUrl,
			"shortUrl": shortUrl,
		})
	})

	r.GET("/tiny/:shortUrl", func(c *gin.Context) {

		shortUrl := c.Param("shortUrl")

		longUrl := GetLongURLFromShortURL(shortUrl)

		c.Redirect(http.StatusMovedPermanently, longUrl)
	})

	err := r.Run()
	if err != nil {
		os.Exit(-1)
	}
}

func getShortURLFromLongURL(longUrl string) string {
	// TODO: Implement
	shortUrl := "x3871hkda"
	return shortUrl
}

func GetLongURLFromShortURL(shortUrl string) string {
	// TODO: Implement this Function
	return "https://www.google.com"
}
