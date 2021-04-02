package core

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

var (
	ClientMap *ClientMapStruct
)

func init() {
	ClientMap = &ClientMapStruct{}
}

// ClientMapStruct 客户端连接对象
type ClientMapStruct struct {
	data sync.Map
}

// Store 保存客户端连接
func (mapStruct *ClientMapStruct) Store(conn *websocket.Conn) {
	client := NewWsClient(conn)
	mapStruct.data.Store(client.conn.RemoteAddr().String(), client)

	go client.Ping(time.Second * 10)
}

// Remove 从map中删除客户端
func (mapStruct *ClientMapStruct) Remove(conn *websocket.Conn) {
	mapStruct.data.Delete(conn.RemoteAddr().String())
}

// SendToAll 向所有连接对象发送消息
func (mapStruct *ClientMapStruct) SendToAll(message string) {
	mapStruct.data.Range(func(key, value interface{}) bool {

		if client, ok := value.(*WsClient); ok {
			if err := client.conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				mapStruct.Remove(client.conn)
				log.Fatalln(err.Error())
			}
		}

		return true
	})
}
