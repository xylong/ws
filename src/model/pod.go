package model

import "fmt"

type PodModel struct {
	PodName  string
	PodImage string
	PodNode  string
}

func MockPodList() []*PodModel {
	return []*PodModel{
		{PodName: "one", PodImage: "nginx:1.18", PodNode: "node1"},
		{PodName: "two", PodImage: "nginx:1.18", PodNode: "node1"},
		{PodName: "three", PodImage: "nginx:1.18", PodNode: "node1"},
	}
}

// ParseAction 解析行为
// 从消息中解析出执行方法
func (model *PodModel) ParseAction(action string) error {
	fmt.Println(action)
	return nil
}
