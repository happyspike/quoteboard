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
