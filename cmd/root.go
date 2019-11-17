package cmd

import (
	"net/http"

	"github.com/triardn/simple-chat-app/handler"
)

func SetupRoutes() {
	http.HandleFunc("/", handler.Homepage)
	http.HandleFunc("/websocket", handler.Websocket)
	http.HandleFunc("/chats", handler.GetAllMessages)
}
