package controllers

type CommentModel struct {
	Captcha Captcha
	ID      uint
	Type    string
}

func NewCommentModel(id uint, t string) CommentModel {
	return CommentModel{
		Captcha: new(CaptchaController).Get(),
		ID:      id,
		Type:    t,
	}
}

type CommentController struct {
}
