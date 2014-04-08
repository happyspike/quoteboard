package main

import (
	"github.com/wkirschbaum/quoteboard/app"
	"html/template"
	"net/http"
	"time"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderQuotePage(w)
	} else if r.Method == "POST" {
		storeQuote(w, r.FormValue("quote"), r.FormValue("author"))
		http.Redirect(w, r, "/", 302)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":4000", nil)
}

func storeQuote(w http.ResponseWriter, content string, author string) {
	quote := app.Quote{
		Content:        content,
		Author:         author,
		Documentor:     "Unknown",
		DocumentedDate: time.Now()}
	makeQuoteStore().Save(quote)
}

var templates = template.Must(template.ParseFiles("public/index.html"))

func renderQuotePage(w http.ResponseWriter) {
	page := app.QuotePage{Quotes: makeQuoteStore().GetAllByDocumentedDateDesc()}

	// false: cache views, true: don't cache views
	devMode := true
	if devMode {
		t, _ := template.ParseFiles("public/index.html")
		t.Execute(w, page)
	} else {
		err := templates.ExecuteTemplate(w, "index.html", page)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func makeQuoteStore() app.QuoteStore {
	return app.QuoteStore{DataFolder: "./data/"}
}
