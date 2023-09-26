package service

import "rest-server/data"

type Service struct {
	dictionary *data.Dictionary
	Words *WordsService
}

func NewService(dictionary *data.Dictionary) *Service {
	return &Service{
		dictionary: dictionary,
		Words: &WordsService{
			dictionary,
		},
	}
}