package main

import (
	"fmt"
	"net/http"

	// Third-party packages
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello, world!")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	http.ListenAndServe("127.0.0.1:8080", router)
}
