package app

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type QuoteStore struct {
	ServerList string
	Database   string
}

func MakeQuoteStore() *QuoteStore {
	return &QuoteStore{
		ServerList: "localhost",
		Database:   "quoteboard"}
}

func (store QuoteStore) Save(quote Quote) {
	session, _ := mgo.Dial(store.ServerList)
	defer session.Close()
	c := session.DB(store.Database).C("quote")
	quote.DocumentedDate = time.Now()
	c.Insert(quote)
}

func (store QuoteStore) GetAllByDocumentedDateDesc() []Quote {
	session, _ := mgo.Dial(store.ServerList)
	defer session.Close()
	c := session.DB(store.Database).C("quote")
	quotes := []Quote{}
	c.Find(bson.M{}).Sort("-documenteddate").All(&quotes)
	return quotes
}
