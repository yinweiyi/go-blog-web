package main

import (
	"blog/bootstrap"
	"blog/config"
	c "blog/vendors/config"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化配置信息
	config.Initialize()
	//初始化数据库
	bootstrap.SetupDB()
	//初始化redis
	bootstrap.SetupRedis()
}

func main() {
	engine := gin.Default()
	//加载路由
	bootstrap.SetupRoute(engine)

	engine.Run(":" + c.GetString("app.port")) // 监听并在 0.0.0.0:8080 上启动服务
}
