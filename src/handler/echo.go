package handler

import (
	"log"
	"net/http"
	"ws/src/core"
)

// Echo ws连接
func Echo(writer http.ResponseWriter, request *http.Request) {
	client, err := core.Upgrader.Upgrade(writer, request, nil) // 协议升级
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		core.ClientMap.Store(client)
	}
}
