//+build generate

package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"

	"github.com/zserge/lorca"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	lorca.Embed("main", "assets.go", "app/build")

	data, err := ioutil.ReadFile("./app/build/index.html")
	check(err)
	b64data := base64.StdEncoding.EncodeToString(data)
	err = ioutil.WriteFile("./app/build/index.html.b64", []byte(b64data), 0644)
	check(err)
}
