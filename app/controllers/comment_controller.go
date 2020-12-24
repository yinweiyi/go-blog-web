package controllers

type Comment struct {
	Captcha Captcha
	ID      uint
	Type    string
}

func NewCommentModel(id uint, types string) Comment {
	return Comment{
		Captcha: new(CaptchaController).Get(),
		ID:      id,
		Type:    types,
	}
}

type CommentController struct {
}
