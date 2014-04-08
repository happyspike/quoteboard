package main

import (
	"github.com/wkirschbaum/quoteboard/app"
	"html/template"
	"net/http"
	"time"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		page := app.QuotePage{Quotes: app.QuoteStore{DataFolder: "./data/"}.GetAllByDocumentedDateDesc()}
		t, _ := template.ParseFiles("public/index.html")
		t.Execute(w, page)
	} else if r.Method == "POST" {
		store := app.QuoteStore{DataFolder: "./data/"}
		quote := app.Quote{
			Content:        r.FormValue("quote"),
			Author:         r.FormValue("author"),
			Documentor:     "Unknown",
			DocumentedDate: time.Now()}
		store.Save(quote)
		http.Redirect(w, r, "/", 302)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":4000", nil)
}
