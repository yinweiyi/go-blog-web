package controllers

import (
	"blog/app/models"
	"blog/app/services"
	configRedis "blog/vendors/redis/config"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (i *ArticleController) Show(ctx *gin.Context) {
	slug := ctx.Param("slug")
	article, err := new(services.ArticleService).GetBySlug(slug)
	i.FailOnError(ctx, err)

	config, err := configRedis.Get()
	i.FailOnError(ctx, err)

	ctx.HTML(200, "article/show.html", gin.H{
		"Config":          config,
		"Sentence":        new(services.SentenceService).GetOne(),
		"Article":         article,
		"MinTags":         models.Shuffle(new(services.TagService).MinTags()),
		"Hots":            new(services.ArticleService).Hots(10),
		"FriendshipLinks": new(services.FriendshipLinkService).Chuck(2),
		"Categories":      new(services.CategoryService).GetAll(),
	})
}
