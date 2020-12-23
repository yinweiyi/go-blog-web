package services

import (
	"blog/app/models"
	"blog/vendors/model"
)

type CategoryService struct {
}

func (CategoryService) GetAll() []models.Category {
	var category []models.Category
	model.DB.Select([]string{"slug", "name"}).Order("`order` asc").Find(&category)
	return category
}

func (CategoryService) GetBySlug(slug string) (models.Category, error) {
	var category models.Category

	err := model.DB.Select([]string{"id", "name"}).First(&category, map[string]interface{}{"slug": slug}).Error
	return category, err
}
