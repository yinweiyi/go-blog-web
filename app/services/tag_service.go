package services

import (
	"blog/app/models"
	"blog/vendors/model"
)

type TagService struct {
}

func (TagService) MinTags() []models.Min {
	var min []models.Min

	model.DB.Model(models.Tag{}).Find(&min)
	return min
}
