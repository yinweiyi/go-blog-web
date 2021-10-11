package services

import (
	"blog/app/models"
	"blog/vendors/model"
)

type AboutService struct {
}

func (a AboutService) All() (abouts []models.About) {
	model.DB.Model(models.About{}).Where("is_enable = ?", 1).Order("`order` asc").Find(&abouts)
	return abouts
}
