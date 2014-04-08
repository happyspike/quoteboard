package app

import "time"

type Quote struct {
	Content        string
	Author         string
	Documentor     string
	DocumentedDate time.Time
}

func (quote Quote) IsValid() bool {
	if len(quote.Content) > 0 &&
		len(quote.Author) > 0 &&
		len(quote.Documentor) > 0 &&
		!quote.DocumentedDate.IsZero() {
		return true
	} else {
		return false
	}
}

type ByDocumentedDateDesc []Quote

func (a ByDocumentedDateDesc) Len() int {
	return len(a)
}

func (a ByDocumentedDateDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByDocumentedDateDesc) Less(i, j int) bool {
	return a[i].DocumentedDate.After(a[j].DocumentedDate)
}
