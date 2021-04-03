package core

import "encoding/json"

// WsMessage websocket消息
type WsMessage struct {
	MsgType int
	MsgData []byte
}

func NewWsMessage(msgType int, msgData []byte) *WsMessage {
	return &WsMessage{MsgType: msgType, MsgData: msgData}
}

// parseToCommand 将消息解析成命令
func (message *WsMessage) parseToCommand() error {
	cmd := &WsCommand{}
	if err := json.Unmarshal(message.MsgData, cmd); err != nil {
		return err
	}

	return cmd.Parse()
}
