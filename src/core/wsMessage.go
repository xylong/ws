package core

// WsMessage websocket消息
type WsMessage struct {
	MsgType int
	MsgData []byte
}

func NewWsMessage(msgType int, msgData []byte) *WsMessage {
	return &WsMessage{MsgType: msgType, MsgData: msgData}
}
