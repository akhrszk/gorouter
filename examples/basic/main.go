package main

import (
	"fmt"
	"net/http"

	"github.com/akhrszk/gorouter"
)

func hello(w http.ResponseWriter, r *http.Request, params gorouter.Params) {
	name := params["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	rt := gorouter.New()
	rt.Get("/hello/:name", hello)
	http.ListenAndServe(":3000", rt)
}
