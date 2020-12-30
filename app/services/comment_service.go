package services

import (
	"blog/app/models"
	"blog/app/models/forms"
	"blog/vendors/helpers"
	"blog/vendors/model"
	"blog/vendors/pagination"
	"errors"
	"net/http"
)

type CommentService struct {
}

var Types = map[string]string{
	"article":   "App\\Models\\Article",
	"about":     "App\\Models\\About",
	"guestbook": "App\\Models\\Guestbook",
}

//Comment 评论
func (c CommentService) Comment(data forms.Comment) (models.Comment, error) {
	commentType, ok := Types[data.Type]
	if !ok {
		return models.Comment{}, errors.New("评论类型不存在")
	}

	var topId = 0

	if data.ParentId > 0 {
		parent, err := c.GetById(data.ParentId)
		if err != nil {
			return models.Comment{}, err
		}
		if parent.TopId > 0 {
			topId = parent.TopId
		} else {
			topId = int(parent.ID)
		}
	}

	comment := models.Comment{
		ParentId:        data.ParentId,
		Content:         data.Content,
		Nickname:        data.Nickname,
		Email:           data.Email,
		CommentableId:   data.ID,
		CommentableType: commentType,
		Ip:              data.IP,
		IsAudited:       1,
		TopId:           topId,
		Avatar:          "/images/avatar.jpg",
	}

	err := model.DB.Create(&comment).Error
	return comment, err
}

//GetById 根据id获取
func (c CommentService) GetById(parentId int) (models.Comment, error) {
	var comment models.Comment
	err := model.DB.First(&comment, parentId).Error
	return comment, err
}

//News 最新评论
func (c CommentService) News() []models.Comment {
	var comments []models.Comment
	model.DB.Where("commentable_type = ? and parent_id = ?", "App\\Models\\Guestbook", 0).Order("id desc").Limit(6).Find(&comments)
	return comments
}

//Count 评论数
func (c CommentService) Count(commentableId uint, commentableType string) int {

	commentType, ok := Types[commentableType]
	if !ok {
		return 0
	}
	where := map[string]interface{}{
		"is_audited":       1,
		"commentable_id":   commentableId,
		"commentable_type": commentType,
	}
	var count int64
	model.DB.Model(models.Comment{}).Where(where).Count(&count)

	return int(count)
}

//GetTree  获取tree
func (c CommentService) GetTree(r *http.Request, perPage, commentableId int, commentableType string) (tree []models.Comment, pagerData pagination.PagerData) {

	commentType, ok := Types[commentableType]
	if !ok {
		return tree, pagerData
	}
	var topComments []models.Comment

	where := map[string]interface{}{
		"top_id":           0,
		"is_audited":       1,
		"commentable_id":   commentableId,
		"commentable_type": commentType,
	}

	db := model.DB.Model(models.Comment{}).Where(where).Order("id desc")
	_pager := pagination.New(r, db, perPage, "", "comment-page")

	pagerData = _pager.Paging()
	if err := _pager.Results(&topComments); err != nil {
		return tree, pagerData
	}

	var topCommentsId []int

	for _, comment := range topComments {
		commentId := int(comment.ID)
		if !helpers.ArrayContain(topCommentsId, commentId) {
			topCommentsId = append(topCommentsId, commentId)
		}
	}

	var childrenComments []models.Comment
	model.DB.Model(models.Comment{}).Where("top_id IN ?", topCommentsId).Where("is_audited = 1").Find(&childrenComments)

	allComment := append(topComments, childrenComments...)

	return c.toTree(allComment, 0), pagerData

}

//toTree  获取tree
func (c CommentService) toTree(allComment []models.Comment, id int) []models.Comment {
	var commentsTree []models.Comment

	for _, comment := range allComment {
		if comment.ParentId == id {
			comment.Children = append(comment.Children, c.toTree(allComment, int(comment.ID))...)
			commentsTree = append(commentsTree, comment)
		}
	}
	return commentsTree
}
