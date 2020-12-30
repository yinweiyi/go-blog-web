package controllers

import (
	"blog/app/models"
	"blog/app/services"
	"blog/vendors/pagination"
	configRedis "blog/vendors/redis/config"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (i *ArticleController) Show(ctx *gin.Context) {
	slug := ctx.Param("slug")
	articleService := new(services.ArticleService)
	article, err := articleService.GetBySlug(slug)
	i.FailOnError(ctx, err)

	commentService := new(services.CommentService)
	commentTree, commentPageData := commentService.GetTree(ctx.Request, 5, int(article.ID), "article")

	paginator := pagination.CreatePaginator(commentPageData, 4)

	articleService.Read(article)
	article.Views += 1
	config, err := configRedis.Get()
	i.FailOnError(ctx, err)
	ctx.HTML(200, "article/show.html", gin.H{
		"Config":          config,
		"Sentence":        new(services.SentenceService).GetOne(),
		"Article":         article,
		"Last":            articleService.Last(article),
		"Next":            articleService.Next(article),
		"MinTags":         models.Shuffle(new(services.TagService).MinTags()),
		"Hots":            articleService.Hots(10),
		"FriendshipLinks": new(services.FriendshipLinkService).Chuck(2),
		"Categories":      new(services.CategoryService).GetAll(),
		"CommentArgs":     NewCommentModel(article.ID, "article"),
		"CommentCount":    commentService.Count(article.ID, "article"),
		"CommentTree":     commentTree,
		"PageLinks":       paginator.Links(),
		"NewestComments":  commentService.News(),
		"NavActive":       "",
	})
}
