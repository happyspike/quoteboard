package app

import (
	"github.com/wkirschbaum/quoteboard/asserts"
	"testing"
	"time"
)

func TestStoreQuoteAsJsonFile(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Date(2014, 1, 2, 3, 4, 5, 6, time.Local)}

	store := Store{}
	store.StoreQuote(quote)

	asserts.AssertFileStringEqual(
		"{\"Content\":\"test content\",\"Author\":\"Test Author\",\"Documentor\":\"Test Documentor\",\"DocumentedDate\":\"2014-01-02T03:04:05.000000006+02:00\"}",
		"../data/2014-01-02T03:04:05.000000006+02:00.quote.json",
		t)
}

func TestStoreQuoteAsJsonFileWithOtherDate(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Date(2015, 2, 3, 4, 5, 6, 7, time.Local)}

	store := Store{}
	store.StoreQuote(quote)

	asserts.AssertFileStringEqual(
		"{\"Content\":\"test content\",\"Author\":\"Test Author\",\"Documentor\":\"Test Documentor\",\"DocumentedDate\":\"2015-02-03T04:05:06.000000007+02:00\"}",
		"../data/2015-02-03T04:05:06.000000007+02:00.quote.json",
		t)

}
