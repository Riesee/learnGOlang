package main

import (
	"log"

	"website-checker/fetch"
)

func main() {
	doc,err := fetch.Get("https://google.com")
	if err != nil {
		log.Fatal(err) // geçici olarak fatal ekledim
	}
	log.Println(doc)
}