package cmd

import (
	"github.com/triardn/simple-chat-app/handler"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", handler.Homepage)
	http.HandleFunc("/websocket", handler.Websocket)
	http.HandleFunc("/chats", handler.GetAllMessages)
}
