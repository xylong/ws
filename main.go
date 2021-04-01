package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		client, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Fatalln(err.Error())
		}

		for {
			if err := client.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
				log.Println(err.Error())
			} else {
				time.Sleep(time.Second * 2)
			}
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err.Error())
	}
}
