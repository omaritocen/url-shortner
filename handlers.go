package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func shortenUrl(c *gin.Context) {

	type RequestBody struct {
		LongUrl string `json:"longUrl" binding:"required"`
	}
	var body RequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortUrl := getShortURLFromLongURL(body.LongUrl)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Tiny url is created",
		"longUrl":  body.LongUrl,
		"shortUrl": shortUrl,
	})
}

func getUrl(c *gin.Context) {

	shortUrl := c.Param("shortUrl")

	longUrl := GetLongURLFromShortURL(shortUrl)

	c.Redirect(http.StatusMovedPermanently, longUrl)
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
