package controllers

import (
	"blog/vendors/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type BaseController struct {
}

// ResponseForSQLError 处理 SQL 错误并返回
func (bc BaseController) FailOnError(ctx *gin.Context, err error) {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 3.1 数据未找到
			ctx.String(http.StatusNotFound, "404 文章未找到")
			panic(fmt.Sprintf("%s", err))
		} else {
			// 3.2 数据库错误
			logger.LogError(err)
			ctx.String(http.StatusInternalServerError, "500 服务器内部错误")
			panic(fmt.Sprintf("%s", err))
		}
	}
}
