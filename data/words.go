package data

import (
	"fmt"

	"gorm.io/gorm"
)

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

func NewWordDictionary(db *gorm.DB) *WordsDictionary {
	return &WordsDictionary{db}
}

func (d *WordsDictionary) GetAll() ([]Word, error) {
	words := make([]Word, 0)
	err := d.db.Find(&words).Error

	return words, err
}

func (d *WordsDictionary) GetOne(id int) (Word, error) {
	word := Word{}
	err := d.db.First(&word, id).Error

	return word, err
}

func (d *WordsDictionary) Delete(id int) (error) {
	err := d.db.Delete(&Word{}, id).Error

	return err
}

func (d *WordsDictionary) DeleteAll() (error) {
	err := d.db.Where("1 = 1").Delete(&Word{}).Error

	return err
}

func (d *WordsDictionary) Update(id int, upd WordProperties) (error) {
	w := Word{}
	err := d.db.Find(&w, id).Error

	if err != nil {
		return err
	}

	if w.ID == 0 {
		return fmt.Errorf("a word with id %d not found", id)
	}

	w.WordProperties = upd

	err = d.db.Save(&w).Error

	return err
}

func (d *WordsDictionary) Add(upd WordProperties) (int, error) {
	word := Word{
		WordProperties: upd,
	}
	err := d.db.Create(&word).Error

	return word.ID, err
}


// TODO: implement DeleteAll method