package models

/*
   关注人信息表(follow)（关系表。用户既是粉丝,又是明星）
   一、可查看某个用户关注的所有人的id,select count(follow_user_id) where(用户id=user_id)
   二、可查看某个用户所有的粉丝id,select count(user_id) where(用户id=follow_user_id)
*/

type Follow struct {
	Model

	UserId       int `gorm:"comment:'用户id';not null" json:"user_id"`
	FollowUserId int `gorm:"comment:'关注用户id';not null;" json:"follow_user_id"`
}
