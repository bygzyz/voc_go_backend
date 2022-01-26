package models

/*
   角色表(role)：不同的角色拥有不同的权限（主要为了实现白名单内的用户才可以发布案例、查看内部情报）
   一、通过 permissions 限定 API 使用权限，初始化可选项为【发布案例、查看内部情报】
*/

type Role struct {
	Model

	Name        string `gorm:"comment:'角色名';size:50;not null;default:''" json:"name"`
	Desc        string `gorm:"comment:'描述信息';size:256;not null;default:''" json:"desc"`
	Permissions string `gorm:"comment:'权限';size:1024;not null;default:''" json:"permissions"`
}
