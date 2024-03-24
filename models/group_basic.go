package models

import "gorm.io/gorm"

// 人员关系
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint // 群主id
	Icon    string
	Desc    string
	Type    int
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
