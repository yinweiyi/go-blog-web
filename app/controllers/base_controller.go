package controllers

import (
	"blog/vendors/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (bc BaseController) Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": message,
		"data":    data,
	})
}

func (bc BaseController) Error(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(200, gin.H{
		"code":    400,
		"message": message,
		"data":    data,
	})
}
