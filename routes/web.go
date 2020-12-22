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
}
