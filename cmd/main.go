package main

import (
	"fmt"
	"net/http"

	"github.com/nimgo/gomux"
)

func main() {

	mux := gomux.New()
	mux.GET("/hello/*watch", flush("Hello!"))
	mux.GET("/helloinline", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello inline!")
	})
}

func flush(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps := gomux.GetMuxParams(r)
		fmt.Fprintf(w, msg+" ...."+ps.ByName("watch"))
	}
}
