package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"ws/src/core"
	"ws/src/handler"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/echo", handler.Echo)
	http.HandleFunc("/all", func(writer http.ResponseWriter, request *http.Request) {
		// msg := request.URL.Query().Get("msg")
		core.ClientMap.SendAllPods()

		if _, err := writer.Write([]byte("ok")); err != nil {
			log.Fatalln(err)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err.Error())
	}
}
