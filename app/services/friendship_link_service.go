package services

import (
	"blog/app/models"
	"blog/vendors/model"
)

type FriendshipLinkService struct {
}

func (FriendshipLinkService) getAll() ([]models.FriendshipLink, error) {
	var friendshipLinks []models.FriendshipLink
	err := model.DB.Where(map[string]interface{}{"is_enable": 1}).Find(&friendshipLinks).Error
	return friendshipLinks, err
}

func (f FriendshipLinkService) Chuck(size int) [][]models.FriendshipLink {
	friendshipLinks, _ := f.getAll()

	var result [][]models.FriendshipLink
	for {
		if len(friendshipLinks) <= size {
			result = append(result, friendshipLinks)
			break
		}
		result = append(result, friendshipLinks[:size])
		friendshipLinks = friendshipLinks[size:]
	}

	return result
}
