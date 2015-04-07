package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddQuote(t *testing.T) {
	quoteboard := MakeQuoteboard()
	quote := MakeQuote()

	quoteboard.AddQuote(quote)
	retrievedQuotes := quoteboard.GetQuotes()

	assert.Len(t, retrievedQuotes, 1)

	assert.Equal(t, 123, 123, "they should be equal")
}

func TestFailingSpec(t *testing.T) {
	assert.Equal(t, 1, 2)
}
