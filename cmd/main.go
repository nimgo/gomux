package main

import (
	"fmt"
	"net/http"

	"github.com/nimgo/nim"
	"github.com/nimgo/nimux"
)

func main() {

	mux := nimux.New()
	mux.GET("/hello/*watch", flush("Hello!"))
	mux.GET("/helloinline", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello inline!")
	})

	auth := nimux.New()
	{
		auth.GET("/auth/boy/:pants", flush("boy"))
		auth.GET("/auth/girl", flush("girl"))
	}

	sub := nim.New()
	sub.WithFunc(middlewareA)
	sub.WithFunc(middlewareB)
	sub.With(auth)

	mux.GET("/auth/*sub", sub.ServeHTTP)

	n := nim.Default()
	n.With(mux)

	nim.Run(n, ":3000")
}

func flush(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, msg)
		ps := nimux.GetHttpParams(r)
		fmt.Println("...." + ps.ByName("watch"))
	}
}

func middlewareA(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[nim.] I am middlewareA")
	//bun := hax.GetBundle(c)
	//bun.Set("valueA", ": from middlewareA")
}

func middlewareB(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[nim.] I am middlewareB")
	//bun := hax.GetBundle(c)
	//bun.Set("valueB", ": from middlewareB")
}
