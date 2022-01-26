package models

import (
	"github.com/golang-module/carbon"
)

// Model 由于gorm提供的 base model 没有 json tag，使用自定义
type Model struct {
	Id        uint                    `gorm:"primaryKey;comment:'自增编号'" json:"id"`
	CreatedAt carbon.ToDateTimeString `gorm:"comment:'创建时间'" json:"created_at"`
	UpdatedAt carbon.ToDateTimeString `gorm:"comment:'更新时间'" json:"updated_at"`
	DeletedAt carbon.ToDateTimeString `gorm:"index:idx_deleted_at;comment:'删除时间(软删除)'" json:"deleted_at"`
	CreatorId int                     `gorm:"comment:'创建人id';default:1" json:"creator_id"`
	// DeletedAt 当 定义该字段时,gorm默认为软删除;普通查询自动忽略；如果需要查询已被软删除的记录加上db.UnScoped() 即可
}
