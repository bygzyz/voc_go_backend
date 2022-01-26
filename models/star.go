package models

// UserArticleStar （用户文章）点赞关系表
type UserArticleStar struct {
	Model
	UserId    int `gorm:"comment:'用户id';not null" json:"user_id"`
	ArticleId int `gorm:"comment:'文章id';not null;" json:"article_id"`
}

// UserCommentStar （用户评论）点赞关系表
type UserCommentStar struct {
	Model
	UserId    int `gorm:"comment:'用户id';not null" json:"user_id"`
	CommentId int `gorm:"comment:'文章id';not null;" json:"comment_id"`
}
