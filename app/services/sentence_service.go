package services

import (
	"blog/app/models"
	"blog/vendors/model"
)

type SentenceService struct {
}

func (SentenceService) GetOne() models.Sentence {
	var sentence models.Sentence
	model.DB.Order("id desc").First(&sentence)
	return sentence
}
