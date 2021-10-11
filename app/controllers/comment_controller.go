package controllers

import (
	"blog/app/models/forms"
	"blog/app/services"
	"blog/vendors/validate"

	"github.com/gin-gonic/gin"
)

type CommentParams struct {
	Captcha Captcha
	ID      uint
	Type    string
}

func NewCommentModel(id uint, types string) CommentParams {
	return CommentParams{
		Captcha: new(CaptchaController).Get(),
		ID:      id,
		Type:    types,
	}
}

type CommentController struct {
	BaseController
}

func (c CommentController) Comment(ctx *gin.Context) {
	var form forms.Comment

	if err := ctx.ShouldBind(&form); err != nil {
		c.FailOnError(ctx, err)
	}
	validator := validate.GetValidator()
	if err := validator.Struct(form); err != nil {
		c.Error(ctx, validate.TranslateOverride(err), nil)
		return
	}
	//验证码验证
	if !new(CaptchaController).Verify(form.CaptchaId, form.Captcha) {
		c.Error(ctx, "验证码错误", nil)
		return
	}
	form.IP = ctx.ClientIP()

	if comment, err := new(services.CommentService).Comment(form); err == nil {
		c.Success(ctx, "评论成功", gin.H{
			"comment_id": comment.ID,
		})
		return
	}
	c.Error(ctx, "评论失败", nil)
}
