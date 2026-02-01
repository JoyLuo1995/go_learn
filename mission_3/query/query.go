//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
package query

import (
	"gorm.io/gorm"
	"go_learn/mission_3/models"
)

func QueryPostsAndCommentsByUserID(db *gorm.DB, userID uint) ([]models.Post, []models.Comment, error) {
	var posts []models.Post
	err :=db.Debug().Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		return nil, nil, err
	}
	var comments []models.Comment
	postIDs := make([]uint, len(posts))
	for i, post := range posts {
		postIDs[i] = post.ID
	}

	err = db.Debug().Where("post_id IN ?", postIDs).Find(&comments).Error
	if err != nil {
		return posts, nil, err
	}
	return posts, comments, err
}