package app

import (
	"fmt"
	"net/http"
)

func Mount(mux *http.ServeMux) {
	fmt.Println("mount")
	mux.HandleFunc("/", index)
}
