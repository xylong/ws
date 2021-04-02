package core

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

var (
	ClientMap *ClientMapStruct
)

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map
}

func (cms *ClientMapStruct) Store(key string, conn *websocket.Conn) {
	cms.data.Store(key, conn)
}

// SendToAll 向所有连接对象发送消息
func (cms *ClientMapStruct) SendToAll(message string) {
	cms.data.Range(func(key, value interface{}) bool {
		if err := value.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Fatalln(err.Error())
		}
		return true
	})
}
