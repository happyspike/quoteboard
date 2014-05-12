package main

import (
	"fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/wkirschbaum/quoteboard/app"
	"time"
)

type QuotePost struct {
	Content string `json: "content" binding:"required"`
	Author  string `json: "author"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Post("/quotes", binding.Json(QuotePost{}), func(quote QuotePost) {
		fmt.Println(quote)
		storeQuote(quote.Content, quote.Author)
	})

	m.Get("/quotes", func(r render.Render) {
		quotes := app.MakeQuoteStore().GetAllByDocumentedDateDesc()
		r.JSON(200, quotes)
	})

	m.Run()
}

func storeQuote(content string, author string) {
	quote := app.Quote{
		Content:        content,
		Author:         author,
		Documentor:     "Unknown",
		DocumentedDate: time.Now()}
	app.MakeQuoteStore().Save(quote)
}
