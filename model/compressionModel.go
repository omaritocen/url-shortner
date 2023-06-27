package model

import "gorm.io/gorm"

type Compression struct {
	gorm.Model
	ShortUrl string `json:"shortUrl"`
	LongUrl  string `json:"longUrl"`
}
