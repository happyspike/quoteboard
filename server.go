package main

import (
	"github.com/wkirschbaum/quoteboard/app"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func main() {
	http.HandleFunc("/assets/", makeHandler(staticHandler))
	http.HandleFunc("/", makeHandler(viewHandler))
	http.HandleFunc("/quotes", makeHandler(quotesHandler))
	http.ListenAndServe(":4000", nil)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderQuotePage(w)
	} else if r.Method == "POST" {
		storeQuote(w, r.FormValue("quote"), r.FormValue("author"))
		http.Redirect(w, r, "/", 302)
	}
}

func quotesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderQuotes(w)
	} else {
		w.WriteHeader(404)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("public" + r.URL.String())
	if err == nil {
		w.WriteHeader(200)
		w.Write(content)
	} else {
		render404Page(w)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		validPathRegex := regexp.MustCompile("^/$|^/assets/\\S+$|^/quotes$")
		matcher := validPathRegex.FindStringSubmatch(r.URL.Path)
		if matcher == nil {
			render404Page(w)
			return
		}
		fn(w, r)
	}
}

var templates = template.Must(template.ParseFiles("public/index.html", "public/404.html"))

func render404Page(w http.ResponseWriter) {
	w.WriteHeader(404)
	renderPage(w, "404", nil)
}

func renderQuotes(w http.ResponseWriter) {
	page := app.QuotePage{Quotes: makeQuoteStore().GetAllByDocumentedDateDesc()}
	renderPage(w, "quotes", page)
}

func renderQuotePage(w http.ResponseWriter) {
	page := app.QuotePage{Quotes: makeQuoteStore().GetAllByDocumentedDateDesc()}
	renderPage(w, "index", page)
}

func renderPage(w http.ResponseWriter, templateName string, object interface{}) {
	devMode := true
	if devMode {
		t, _ := template.ParseFiles("public/" + templateName + ".html")
		err := t.Execute(w, object)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := templates.ExecuteTemplate(w, templateName+".html", object)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func storeQuote(w http.ResponseWriter, content string, author string) {
	quote := app.Quote{
		Content:        content,
		Author:         author,
		Documentor:     "Unknown",
		DocumentedDate: time.Now()}
	makeQuoteStore().Save(quote)
}

func makeQuoteStore() app.QuoteStore {
	return app.QuoteStore{DataFolder: "./data/"}
}
