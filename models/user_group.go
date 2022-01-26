package models

type UserGroup struct {
	Model

	Name       string `gorm:"comment:'角色名';size:50;not null;default:''" json:"name"`
	Desc       string `gorm:"comment:'描述信息';size:256;not null;default:''" json:"desc"`
	Permission string `gorm:"comment:'权限';size:1024;not null;default:''" json:"permission"`
	IsActivate bool   `gorm:"comment:'是否激活';not null;default:true" json:"is_activate"`
}
