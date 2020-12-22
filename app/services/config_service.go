package services

import (
	"blog/app/models"
	"blog/vendors/model"
)

type ConfigService struct {
}

func (ConfigService) GetOne() (models.Config, error) {
	var config models.Config
	err := model.DB.First(&config).Error
	return config, err
}
