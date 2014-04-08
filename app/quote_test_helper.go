package app

import (
	"time"
)

func createQuoteNow(content string) Quote {
	return Quote{
		Content:        content,
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Now()}
}
