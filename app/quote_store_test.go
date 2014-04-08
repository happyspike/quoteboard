package app

import (
	"github.com/wkirschbaum/quoteboard/asserts"
	"os"
	"testing"
	"time"
)

func TestStoreQuoteMarchallFilenameDate1(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Date(2015, 2, 3, 4, 5, 6, 7, time.Local)}

	store := QuoteStore{DataFolder: "../test_data/"}
	filename := store.MarshalFilename(quote)
	asserts.AssertStringEqual("../test_data/2015-02-03T04:05:06.000000007+02:00.quote.json", filename, t)
}

func TestStoreQuoteMarchallFilenameDate2(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Date(2014, 1, 2, 3, 4, 5, 6, time.Local)}

	store := QuoteStore{DataFolder: "../test_data/"}
	filename := store.MarshalFilename(quote)
	asserts.AssertStringEqual("../test_data/2014-01-02T03:04:05.000000006+02:00.quote.json", filename, t)
}

func TestQuoteStoreSaveAsJsonFile(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Date(2014, 1, 2, 3, 4, 5, 6, time.Local)}

	store := QuoteStore{DataFolder: "../test_data/"}
	store.Save(quote)

	asserts.AssertFileStringEqual(
		"{\"Content\":\"test content\",\"Author\":\"Test Author\",\"Documentor\":\"Test Documentor\",\"DocumentedDate\":\"2014-01-02T03:04:05.000000006+02:00\"}",
		store.MarshalFilename(quote),
		t)
	cleanData()
}

func TestQuoteStoreSaveAsJsonFileWithOtherDate(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Date(2015, 2, 3, 4, 5, 6, 7, time.Local)}

	store := QuoteStore{DataFolder: "../test_data/"}
	store.Save(quote)

	asserts.AssertFileStringEqual(
		"{\"Content\":\"test content\",\"Author\":\"Test Author\",\"Documentor\":\"Test Documentor\",\"DocumentedDate\":\"2015-02-03T04:05:06.000000007+02:00\"}",
		store.MarshalFilename(quote),
		t)
	cleanData()
}

func TestStoreGetAllQuotes(t *testing.T) {
	quote1 := createQuoteNow("test quote 1")
	quote2 := createQuoteNow("test quote 2")
	store := QuoteStore{DataFolder: "../test_data/"}
	store.Save(quote1)
	store.Save(quote2)

	quotes := store.GetAll()

	if quotes[0].Content != quote1.Content {
		t.Fail()
	}

	if quotes[1].Content != quote2.Content {
		t.Fail()
	}
	cleanData()
}

func TestGetAllByDocumentedDateDesc(t *testing.T) {
	quote1 := createQuoteNow("quote1")
	quote2 := createQuoteNow("quote2")
	quote3 := createQuoteNow("quote3")
	store := QuoteStore{DataFolder: "../test_data/"}
	store.Save(quote1)
	store.Save(quote3)
	store.Save(quote2)

	sorted_quotes := store.GetAllByDocumentedDateDesc()
	asserts.AssertStringEqual("quote3", sorted_quotes[0].Content, t)
	asserts.AssertStringEqual("quote2", sorted_quotes[1].Content, t)
	asserts.AssertStringEqual("quote1", sorted_quotes[2].Content, t)

	cleanData()
}

func cleanData() {
	os.RemoveAll("../test_data/")
}
