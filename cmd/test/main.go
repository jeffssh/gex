package main

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jeffssh/gex"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type PayLoad struct {
	Content string
}

func main() {
	// consider passign in pipes
	// https://golang.org/pkg/io/#Pipe
	lr, lw := io.Pipe()
	rr, rw := io.Pipe()
	go func() {
		for {
			lw.Write([]byte("Testing Gex!"))
			time.Sleep(5 * time.Second)
		}
	}()

	go func() {
		buff := make([]byte, 2<<16)
		for {
			bytesRead, err := rr.Read(buff)
			data := buff[:bytesRead]
			check(err)
			fmt.Println(string(data))

		}
	}()
	g, err := gex.New(lr, rw, 2<<16)
	check(err)
	if g == nil {
		fmt.Println("g is nil")
	}
	g.Serve()

}
