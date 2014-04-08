package app

import (
	"io/ioutil"
	"os"
)

type Store struct {
}

func (store Store) StoreQuote(quote Quote) {
	ioutil.WriteFile("../data/2014-01-02.03:04:05.000000006.quote.json", []byte("{\"Content\":\"test content\"}"), os.ModePerm)
}
