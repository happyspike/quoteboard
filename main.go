package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	n := negroni.Classic()
	n.UseHandler(mux)
	StartHttpServer(n)
}

func StartHttpServer(handler http.Handler) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	addr := ":" + port
	log.Println("Starting server on " + addr)
	http.ListenAndServe(addr, handler)
}
