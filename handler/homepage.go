package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("view/index.html")
	if err != nil {
		http.Error(w, "Could not open requested file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", content)
}
