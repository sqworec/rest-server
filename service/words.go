package service

import (
	"fmt"
	"rest-server/data"
)

type WordsService struct {
	dictionary *data.Dictionary
}

func (d *WordsService) GetAll() ([]data.Word, error) {
	words, err := d.dictionary.Words.GetAll()

	return words, err
}

func (d *WordsService) GetOne(id int) (data.Word, error) {
	word, err := d.dictionary.Words.GetOne(id)

	return word, err
}

func (d *WordsService) Delete(id int) (error) {
	err := d.dictionary.Words.Delete(id)

	return err
}

func (d *WordsService) DeleteAll() (error) {
	err := d.dictionary.Words.DeleteAll()

	return err
}

func (d *WordsService) Update(id int, upd data.WordProperties) (error) {

	if upd.EnglishWord == "" || upd.RussianWord == "" {
		return fmt.Errorf("Service.Update: russian or english word must not be empty")
	}

	err := d.dictionary.Words.Update(id, upd)

	return err 
}

func (d *WordsService) Add(upd data.WordProperties) (int, error) {

	if upd.EnglishWord == "" || upd.RussianWord == "" {
		return 0, fmt.Errorf("Service.Update: russian or english word must not be empty")
	}

	id, err := d.dictionary.Words.Add(upd)

	return id, err 
}