package main

import (
	"log"

	"github.com/jeffssh/gex"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	gex.New(nil, nil)
	log.Println("Done with gex.New()")
}
