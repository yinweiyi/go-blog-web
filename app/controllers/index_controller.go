package controllers

import (
	"blog/app/models"
	"blog/app/services"
	"blog/vendors/model"
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
	cate, err := new(services.CategoryService).GetBySlug(slug)
	if err != nil {
		i.rendor(ctx, []models.Article{}, pagination.PagerData{})
	} else {
		articles, pagerData, err := new(services.ArticleService).GetAll(ctx.Request, 5, map[string]interface{}{"category_id": cate.ID})
		i.FailOnError(ctx, err)
		i.rendor(ctx, articles, pagerData)
	}

}

func (i *IndexController) Tag(ctx *gin.Context) {
	slug := ctx.Param("slug")
	tagService := new(services.TagService)

	tag, err := tagService.GetBySlug(slug)
	if err != nil {
		i.rendor(ctx, []models.Article{}, pagination.PagerData{})
	} else {
		articles, pagerData, err := tagService.GetArticlesByTag(ctx.Request, tag, 5)
		i.FailOnError(ctx, err)
		i.rendor(ctx, articles, pagerData)
	}
}

func (i *IndexController) About(ctx *gin.Context) {
	config, err := configRedis.Get()
	i.FailOnError(ctx, err)
	abouts := new(services.AboutService).All()
	commentService := new(services.CommentService)
	var commentTree []models.Comment
	var commentPageData pagination.PagerData
	var aboutId uint

	if len(abouts) > 0 {
		aboutId = abouts[0].ID
		commentTree, commentPageData = commentService.GetTree(ctx.Request, 5, int(aboutId), "about")
	}

	paginator := pagination.CreatePaginator(commentPageData, 4)

	ctx.HTML(200, "index/about.html", gin.H{
		"Config":         config,
		"Abouts":         abouts,
		"MinTags":        models.Shuffle(new(services.TagService).MinTags()),
		"Hots":           new(services.ArticleService).Hots(10),
		"CommentArgs":    NewCommentModel(aboutId, "about"),
		"CommentCount":   commentService.Count(aboutId, "about"),
		"CommentTree":    commentTree,
		"PageLinks":      paginator.Links(),
		"Categories":     new(services.CategoryService).GetAll(),
		"NewestComments": commentService.News(),
		"NavActive":      "about",
	})
}

func (i *IndexController) Guestbook(ctx *gin.Context) {
	config, err := configRedis.Get()
	i.FailOnError(ctx, err)
	var guestbook models.Guestbook
	model.DB.First(&guestbook)

	commentService := new(services.CommentService)
	var commentTree []models.Comment
	var commentPageData pagination.PagerData
	var guestbookId uint

	if guestbook.ID > 0 {
		guestbookId = guestbook.ID
		commentTree, commentPageData = commentService.GetTree(ctx.Request, 5, int(guestbookId), "guestbook")
	}

	paginator := pagination.CreatePaginator(commentPageData, 4)

	ctx.HTML(200, "index/guestbook.html", gin.H{
		"Config":         config,
		"Guestbook":      guestbook,
		"MinTags":        models.Shuffle(new(services.TagService).MinTags()),
		"Hots":           new(services.ArticleService).Hots(10),
		"CommentArgs":    NewCommentModel(guestbookId, "guestbook"),
		"CommentCount":   commentService.Count(guestbookId, "guestbook"),
		"CommentTree":    commentTree,
		"PageLinks":      paginator.Links(),
		"Categories":     new(services.CategoryService).GetAll(),
		"NewestComments": commentService.News(),
		"NavActive":      "guestbook",
	})
}

func (i *IndexController) rendor(ctx *gin.Context, articles []models.Article, pagerData pagination.PagerData) {

	config, err := configRedis.Get()
	i.FailOnError(ctx, err)

	paginator := pagination.CreatePaginator(pagerData, 4)
	ctx.HTML(200, "index/index.html", gin.H{
		"Config":          config,
		"Sentence":        new(services.SentenceService).GetOne(),
		"Articles":        articles,
		"PageLinks":       paginator.Links(),
		"MinTags":         models.Shuffle(new(services.TagService).MinTags()),
		"Hots":            new(services.ArticleService).Hots(10),
		"FriendshipLinks": new(services.FriendshipLinkService).Chuck(2),
		"Categories":      new(services.CategoryService).GetAll(),
		"NewestComments":  new(services.CommentService).News(),
		"NavActive":       "/",
	})
}
