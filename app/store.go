package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Store struct {
}

func (store Store) StoreQuote(quote Quote) {
	store_json, _ := json.Marshal(quote)
	filename := getFilename(quote)
	ioutil.WriteFile(filename, store_json, os.ModePerm)
}

func getFilename(quote Quote) string {
	date_bytes, _ := quote.DocumentedDate.MarshalJSON()
	date_string := strings.Replace(string(date_bytes), "\"", "", -1)
	return "../data/" + date_string + ".quote.json"
}
