package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// WsClient ws客户端
type WsClient struct {
	conn     *websocket.Conn
	readChan chan *WsMessage // 读队列
	exitChan chan struct{}   // 失败退出队列
}

// NewWsClient 创建websocket客户端
func NewWsClient(conn *websocket.Conn) *WsClient {
	return &WsClient{
		conn:     conn,
		readChan: make(chan *WsMessage),
		exitChan: make(chan struct{}),
	}
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

// readLoop 循环读
func (client *WsClient) readLoop() {
	for {
		if msgType, bytes, err := client.conn.ReadMessage(); err != nil {
			_ = client.conn.Close()
			ClientMap.Remove(client.conn)
			client.exitChan <- struct{}{}
			break
		} else {
			client.readChan <- NewWsMessage(msgType, bytes)
		}
	}
}

// HandleMessage 处理消息
func (client *WsClient) HandleMessage() {
loop:
	for {
		select {
		case msg := <-client.readChan:
			fmt.Println(string(msg.MsgData))
		case <-client.exitChan:
			log.Println("已关闭")
			break loop
		}
	}
}
