package data

import "gorm.io/gorm"

type WordProperties struct {
	RussianWord string `json:"ru"`
	EnglishWord string `json:"en"`
}

type Word struct {
	WordProperties
	ID int `json:"id"`
}

type WordsDictionary struct {
	db *gorm.DB
}

