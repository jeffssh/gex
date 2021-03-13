//+build generate

package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("./app/dist/build.html")
	check(err)
	b64data := base64.StdEncoding.EncodeToString(data)
	fmt.Println(b64data)
	err = ioutil.WriteFile("./app/dist/build.html.b64", []byte(b64data), 0644)
	check(err)
}
