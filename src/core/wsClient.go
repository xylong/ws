package core

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// WsClient ws客户端
type WsClient struct {
	conn *websocket.Conn
}

// NewWsClient 创建websocket客户端
func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{conn: conn}
}

// Ping
// wait 心跳间隔
func (client *WsClient) Ping(wait time.Duration) {
	for {
		time.Sleep(wait)
		// 发心跳
		if err := client.conn.WriteMessage(websocket.TextMessage, []byte("ping")); err != nil {
			ClientMap.Remove(client.conn)
			log.Fatalln(err.Error())
			return
		}
	}
}
