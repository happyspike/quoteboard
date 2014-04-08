package app

import (
	"fmt"
	"io/ioutil"
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

	content, err := ioutil.ReadFile("../data/2014-01-02.03:04:05.000000006.quote.json")
	if err != nil {
		fmt.Println(err.Error())
		t.Fail()
	} else {
		fmt.Println(string(content))
		if string(content) != "{\"Content\":\"test content\"}" {
			t.Fail()
		}
	}

}
