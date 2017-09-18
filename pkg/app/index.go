package app

import (
	"fmt"
	"net/http"

	"github.com/Chayanon/scweb/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	view.Index(w, nil)
	fmt.Println("index")
}
