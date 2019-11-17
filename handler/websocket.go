package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var connections = make([]*WebSocketConnection, 0)

var TempMessages []string

type WebSocketConnection struct {
	*websocket.Conn
}

type SocketPayload struct {
	Message string
}

type SocketResponse struct {
	Type    string
	Message string
}

func Websocket(w http.ResponseWriter, r *http.Request) {
	currentConnection, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
	if err != nil {
		http.Error(w, "could not open websocket connection", http.StatusBadRequest)
	}

	currentConn := WebSocketConnection{Conn: currentConnection}
	connections = append(connections, &currentConn)

	go handleIO(&currentConn, connections)
}

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				return
			}

			log.Println("ERROR", err.Error())
			continue
		}

		broadcastMessage(currentConn, payload.Message)

		// add message to temp variable
		TempMessages = append(TempMessages, payload.Message)
	}
}

func broadcastMessage(currentConn *WebSocketConnection, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}

		eachConn.WriteJSON(SocketResponse{
			Message: message,
		})
	}
}
