package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromId   uint   // 发送着id
	TargetId uint   // 接受者id
	Type     string // 消息类型 ：群聊、私聊、广播
	Media    int    // 消息类型：语音、文字、音视频
	Content  string //消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int // 其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}
