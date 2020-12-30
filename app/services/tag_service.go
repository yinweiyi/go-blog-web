package services

import (
	"blog/app/models"
	"blog/vendors/model"
	"blog/vendors/pagination"
	"net/http"
)

type TagService struct {
}

func (TagService) MinTags() []models.Min {
	var min []models.Min

	model.DB.Model(models.Tag{}).Find(&min)
	return min
}

//GetBySlug 根据slug获取tag
func (TagService) GetBySlug(slug string) (models.Tag, error) {
	var tag models.Tag

	err := model.DB.Select([]string{"id", "name"}).First(&tag, map[string]interface{}{"slug": slug}).Error
	return tag, err
}

func (TagService) GetArticlesByTag(ctx *http.Request, tag models.Tag, perPage int) ([]models.Article, pagination.PagerData, error) {
	db := model.DB.Model(&tag).Preload("Tags").Order("is_top desc, `order`")

	_pager := pagination.New(ctx, db, perPage, "Articles", "comment-page")

	// 2. 获取视图数据
	PagerData := _pager.Paging()

	// 3. 获取数据
	var articles []models.Article
	err := _pager.Results(&articles)
	return articles, PagerData, err
}
