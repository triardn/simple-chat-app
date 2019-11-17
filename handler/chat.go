package handler

import (
	"encoding/json"
	"net/http"
)

type ChatPayload struct {
	Message string `json:"message"`
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if TempMessages == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("No message recorded"))
		return
	}

	resp, err := json.Marshal(TempMessages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error when marshaling data"))
		return
	}

	w.Write(resp)
}
