package models

// Record 统计用户行为
type Record struct {
	Model

	UserId                   int    `gorm:"comment:'用户id';not null;" json:"user_id"`
	SearchCaseHistory        string `gorm:"comment:'案例搜索历史';size:1024;not null;default:''" json:"search_case_history"`
	SearchInformationHistory string `gorm:"comment:'情报搜索历史';size:1024;not null;default:''" json:"search_information_history"`
	LastViewedArticleIds     string `gorm:"comment:'文章浏览历史';size:1024;not null;default:''" json:"last_viewed_article_ids"`
}
