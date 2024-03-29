package routes

import (
	"blog/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoute(engine *gin.Engine) {
	//首页
	index := new(controllers.IndexController)
	engine.GET("/", index.Index)
	engine.GET("/category/:slug", index.Category)
	engine.GET("/tag/:slug", index.Tag)
	engine.GET("/about", index.About)
	engine.GET("/guestbook", index.Guestbook)

	//文章页
	article := new(controllers.ArticleController)
	engine.GET("/articles/:slug", article.Show)

	//验证码
	captcha := new(controllers.CaptchaController)
	engine.GET("/captcha/:captchaId", captcha.Captcha)

	//评论
	comment := new(controllers.CommentController)
	engine.POST("/comment", comment.Comment)

}
