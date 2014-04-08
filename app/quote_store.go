package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type QuoteStore struct {
}

func (store QuoteStore) Save(quote Quote) {
	store_json, _ := json.Marshal(quote)
	filename := store.MarshalFilename(quote)
	os.Mkdir("../data/", os.ModePerm)
	ioutil.WriteFile(filename, store_json, os.ModePerm)
}

func (store QuoteStore) GetAll() []Quote {
	files, _ := ioutil.ReadDir("../data/")
	var quotes = []Quote{}

	for _, fileinfo := range files {
		quote := Quote{}
		filedata, _ := ioutil.ReadFile("../data/" + fileinfo.Name())
		json.Unmarshal(filedata, &quote)
		quotes = append(quotes, quote)
	}
	return quotes
}

func (store QuoteStore) GetAllByDocumentedDateDesc() []Quote {
	quotes := store.GetAll()
	sort.Sort(ByDocumentedDateDesc(quotes))
	return quotes
}

func (store QuoteStore) MarshalFilename(quote Quote) string {
	date_bytes, _ := quote.DocumentedDate.MarshalJSON()
	date_string := strings.Replace(string(date_bytes), "\"", "", -1)
	return "../data/" + date_string + ".quote.json"
}
