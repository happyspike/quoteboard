package app

import (
	"testing"
	"time"
)

func TestQuoteValid(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Now()}

	if quote.IsValid() == false {
		t.Fail()
	}
}

func TestQuoteWithoutContentInvalid(t *testing.T) {
	quote := Quote{
		Author:         "Test Author",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Now()}

	if quote.IsValid() == true {
		t.Fail()
	}
}

func TestQuoteWithoutAuthorInvalid(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Documentor:     "Test Documentor",
		DocumentedDate: time.Now()}

	if quote.IsValid() == true {
		t.Fail()
	}
}

func TestQuoteWithoutDocumentorInvalid(t *testing.T) {
	quote := Quote{
		Content:        "test content",
		Author:         "Test Author",
		DocumentedDate: time.Now()}

	if quote.IsValid() == true {
		t.Fail()
	}
}

func TestQuoteWithoutDocumentedDateInvalid(t *testing.T) {
	quote := Quote{
		Content:    "test content",
		Author:     "Test Author",
		Documentor: "Test Documentor"}

	if quote.IsValid() == true {
		t.Fail()
	}
}
