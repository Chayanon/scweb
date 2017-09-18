package main

import (
	"fmt"
	"net/http"

	"github.com/Chayanon/scweb/pkg/app"
)

const port = ":8080"

func main() {
	fmt.Println("Main")
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}
