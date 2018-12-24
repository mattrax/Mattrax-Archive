package main

import (
	"log"
	"os"

	"github.com/mattrax/Mattrax/pkg/soap"
)

type request struct {
	Head struct {
		MessageID string
	}
	Body struct {
	}
}

func main() {
	file, err := os.Open("./pkg/soap/example/req.soap")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var payload request
	soap.Parse(file, payload)
	log.Println(payload)
}
