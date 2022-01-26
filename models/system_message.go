package models

type SystemMessage struct {
	Model
	Title    string `gorm:"comment:'标题';size:256;not null;default:''" json:"title"`
	Desc     string `gorm:"comment:'描述信息';size:256;not null;default:''" json:"desc"`
	Content  string `gorm:"comment:'内容';size:9999;not null;default:''" json:"content"`
	ImageUrl string `gorm:"comment:'图片链接地址';size:256;not null;default:''" json:"image_url"`
	//ViewUserIds    string `gorm:"comment:头像链接地址;size:256;not null;default:''" json:"view_user_ids"`
}
