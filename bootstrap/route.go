package bootstrap

import (
	"blog/routes"
	"github.com/gin-gonic/gin"
)

// SetupRoute 路由初始化
func SetupRoute(engine *gin.Engine) {
	//加载公共文件以及模板
	routes.RegisterCommonFile(engine)
	//加载前台路由
	routes.RegisterWebRoute(engine)
}
