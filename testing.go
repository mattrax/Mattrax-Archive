package main

import (
	"errors"
	"log"
)

// TODO: Look At "Model" SQL Library For Go

type Testing interface {
	Get(name string) (string, error)
	Set(name string) error
}

/*
  db.Query(struct, "SELECT * FROM repo WHERE $1", name)
*/

func (t Testing) Get(name string) (string, error) {
	return "Hello", errors.New("Testing")
}

func main() {
	log.Println("Hello World")

	test := Testing.Get("Hello World")
}

/*
  type Testing struct {
  TestStr string
}

var d = schema.NewDecoder()
d.Decode(&opt, r.URL().Query())

*/

// Have A Look At https://github.com/sourcegraph/appmon -> Go lang Webserver Insights In Postrges Database

//TODO: Download + Docs
//  DB Client
//  Plist Parsing/Generation Library
