package main

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	n := negroni.Classic()
	n.UseHandler(mux)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	n.Run(":" + port)
}
