package main

import (
	"fmt"
	"github.com/triardn/simple-chat-app/cmd"
	"net/http"
)

func main() {
	cmd.SetupRoutes()

	fmt.Println("server starting at :1234")
	http.ListenAndServe(":1234", nil)
}
