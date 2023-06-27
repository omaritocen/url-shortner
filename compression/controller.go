package compression

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func ShortenUrl(c *gin.Context) {

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

	if compression := getCompressionByLongURL(body.LongUrl); compression != nil {
		c.JSON(http.StatusOK, gin.H{
			"url":      compression.LongUrl,
			"shortUrl": compression.ShortUrl,
		})
		return
	}

	shortUrl := generateShortURL()
	compression := addCompression(shortUrl, body.LongUrl)

	// Return new shortUrl
	c.JSON(http.StatusCreated, gin.H{
		"url":      compression.LongUrl,
		"shortUrl": compression.ShortUrl,
	})
}

func GetUrl(c *gin.Context) {

	shortUrl := c.Param("shortUrl")

	compression := getCompressionByShortURL(shortUrl)
	if compression == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "This id does not exist",
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, compression.LongUrl)
}

func generateShortURL() string {
	uid := uuid.New()
	base62id := EncodeBase62(int64(uid.ID()))
	return base62id
}
