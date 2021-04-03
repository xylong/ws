package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"ws/src/model"
)

const (
	NewPod = 101 // 新增商品
)

var (
	Command = map[int]model.IModel{}
)

func init() {
	Command[NewPod] = (*model.PodModel)(nil) // 反射
}

// WsCommand 命令
type WsCommand struct {
	CmdType   int
	CmdData   map[string]interface{}
	CmdAction string
}

// Parse 执行解析
func (command *WsCommand) Parse() error {
	if v, ok := Command[command.CmdType]; ok {
		// 通过反射获取对象
		obj := reflect.New(reflect.TypeOf(v).Elem()).Interface()
		// 通过json方式，将map转成struct
		bytes, err := json.Marshal(command.CmdData)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(bytes, obj); err != nil {
			return err
		}

		return obj.(model.IModel).ParseAction(command.CmdAction)
	}

	return fmt.Errorf("error command")
}
