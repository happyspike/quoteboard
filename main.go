package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/happyspike/bladerunner"
)

func main() {
	mux := http.NewServeMux()
	n := negroni.Classic()
	n.UseHandler(mux)
	bladerunner.Run(n)
}
