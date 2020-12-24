package controllers

type Comment struct {
	Captcha Captcha
	ID      uint
	Type    string
}

func NewCommentModel(id uint, t string) Comment {
	return Comment{
		Captcha: new(CaptchaController).Get(),
		ID:      id,
		Type:    t,
	}
}

type CommentController struct {
}
