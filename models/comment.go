package models

type Comment struct {
	Model

	ParentId    int    `gorm:"comment:'父评论id';not null;default:0;" json:"parent_id"` // 父评论id为0代表当前记录父评论; 当父评论id不为0,代表当前记录为子评论
	ArticleId   int    `gorm:"comment:'文章id';not null;" json:"article_id"`
	UserId      int    `gorm:"comment:'评论人id';not null;" json:"user_id"`
	ReplyUserId int    `gorm:"comment:'被回复人id';not null;default:0" json:"reply_user_id"`
	Content     string `gorm:"comment:'评论内容';size:255;not null;default:''" json:"content"`
}
