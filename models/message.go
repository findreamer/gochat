package models

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId   uint   // 发送着id
	TargetId uint   // 接受者id
	Type     string // 发送类型 ：群聊、私聊、广播
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

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	// 检验token
	// token := query.Get("token")

	query := request.URL.Query()
	userId := query.Get("userId")
	msyType := query.Get("type")
	targetId := query.Get("targetId")
	content := query.Get("content")
	isvalida := true // checkToken()

	conn, err := (&websocket.Upgrader{
		// token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建Node
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

}
