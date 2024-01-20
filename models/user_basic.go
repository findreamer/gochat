package models

import (
	"fmt"
	"gochat/utils"
	"time"

	"gorm.io/gorm"
)

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
	LoginTime *time.Time
	// 下线时间
	LoginOutTime *time.Time
	// 心跳时间
	HeartbeatTime *time.Time

	IsLogout   bool
	DeviceInfo string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
