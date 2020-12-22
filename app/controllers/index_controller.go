package controllers

import (
	"blog/app/models"
	"blog/app/services"
	"blog/vendors/pagination"
	configRedis "blog/vendors/redis/config"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

func (i *IndexController) Index(ctx *gin.Context) {
	var where map[string]interface{}
	articles, pagerData, err := new(services.ArticleService).GetAll(ctx.Request, 5, where)
	i.FailOnError(ctx, err)

	i.rendor(ctx, articles, pagerData)
}

func (i *IndexController) Category(ctx *gin.Context) {
	slug := ctx.Param("slug")
	cate := new(services.CategoryService).GetBySlug(slug)

	articles, pagerData, err := new(services.ArticleService).GetAll(ctx.Request, 5, map[string]interface{}{"category_id": cate.ID})
	i.FailOnError(ctx, err)
	i.rendor(ctx, articles, pagerData)
}

func (i *IndexController) tag(ctx *gin.Context) {
	//slug := ctx.Param("slug")
}

func (i *IndexController) rendor(ctx *gin.Context, articles []models.Article, pagerData pagination.PagerData) {

	config, err := configRedis.Get()
	i.FailOnError(ctx, err)

	ctx.HTML(200, "index/index.html", gin.H{
		"Config":          config,
		"Sentence":        new(services.SentenceService).GetOne(),
		"Articles":        articles,
		"PagerData":       pagerData,
		"MinTags":         models.Shuffle(new(services.TagService).MinTags()),
		"Hots":            new(services.ArticleService).Hots(10),
		"FriendshipLinks": new(services.FriendshipLinkService).Chuck(2),
		"Categories":      new(services.CategoryService).GetAll(),
	})
}
