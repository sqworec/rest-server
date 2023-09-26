package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConfig struct {
	Path string
}

type Dictionary struct {
	db    *gorm.DB
	Words *WordsDictionary
}

func (d *Dictionary) GetDB() *gorm.DB {
	return d.db
}

func NewDictionary(config DBConfig) *Dictionary {
	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Word{})

	dictionary := &Dictionary{
		db: db,
		Words: NewWordDictionary(db),
	}

	return dictionary
}