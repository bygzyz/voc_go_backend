package service

import "gorm.io/gorm"

type MysqlService struct {
	tx *gorm.DB // 事务对象
	db *gorm.DB // 无事务对象
}
