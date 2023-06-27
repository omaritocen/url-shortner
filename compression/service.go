package compression

import (
	"url-shortner/db"
	"url-shortner/models"
)

func addCompression(shortUrl string, longUrl string) *models.Compression {
	// Store in DB
	db := db.GetDB()

	compression := &models.Compression{
		ShortUrl: shortUrl,
		LongUrl:  longUrl,
	}

	db.Create(compression)

	return compression
}

func getCompressionByLongURL(longUrl string) *models.Compression {
	db := db.GetDB()

	var compression *models.Compression
	if err := db.Where("long_url = ?", longUrl).First(&compression).Error; err != nil {
		return nil
	}

	return compression
}

func getCompressionByShortURL(shortUrl string) *models.Compression {
	db := db.GetDB()

	var compression *models.Compression
	if err := db.Where("short_url = ?", shortUrl).First(&compression).Error; err != nil {
		return nil
	}

	return compression
}
