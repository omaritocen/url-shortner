package models

import "gorm.io/gorm"

type Compression struct {
	gorm.Model
	ShortUrl string `json:"shortUrl" gorm:"primaryKey"`
	LongUrl  string `json:"longUrl"`
}
