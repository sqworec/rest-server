package data

import "gorm.io/gorm"

type DBConfig struct {
	Path string
}

type Dictionary struct {
	db    *gorm.DB
	Words *WordsDictionary
}