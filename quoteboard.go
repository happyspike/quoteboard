package main

// Quoteboard ...
type Quoteboard struct {
	quotes []Quote
}

// MakeQuoteboard ...
func MakeQuoteboard() Quoteboard {
	return Quoteboard{}
}

// AddQuote ...
func (quoteboard *Quoteboard) AddQuote(quote Quote) {
	quoteboard.quotes = append(quoteboard.quotes, quote)
}

// GetQuotes ...
func (quoteboard Quoteboard) GetQuotes() []Quote {
	return quoteboard.quotes
}
