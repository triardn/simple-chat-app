package main

import (
	"fmt"
	"net/http"

	"github.com/triardn/simple-chat-app/cmd"
)

func main() {
	cmd.SetupRoutes()

	fmt.Println("server starting at :1234")
	http.ListenAndServe(":1234", nil)
}
