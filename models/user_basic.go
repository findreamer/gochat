package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name     string
	PassWord string
	Phone    string
	Email    string
	Identity string
	// 设备ip
	ClientIp   string
	clientPort string

	// 登陆时间
	LoginTime uint64
	// 下线时间
	LogoutTime uint64
	// 心跳时间
	HeartbeatTime uint64

	IsLogout   bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
