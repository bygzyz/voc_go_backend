package models

// Collection 收藏表
type Collection struct {
	Model

	UserId    int `gorm:"comment:'用户id';not null;" json:"user_id"`
	ArticleId int `gorm:"comment:'文章id';not null;" json:"article_id"`
}
