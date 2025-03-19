package models

import "time"

type URL struct {
	Id         uint      `json:"id" gorm:"primaryKey"`
	OrginalUrl string    `json:"original-url"`
	ShortenUrl string    `json:"shorten-url"`
	CreatedAt  time.Time `json:"created-at"`
}

type GetUrl struct {
	Url string `json:"url"`
}
