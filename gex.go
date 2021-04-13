package gex

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Gex struct {
	r              *io.Reader
	w              *io.Writer
	websocketConn  *websocket.Conn
	websocketMutex *sync.Mutex
	html           string
	upgrader       *websocket.Upgrader
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doRead() {}

func (g *Gex) serveApp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string(g.html))
}

func (g *Gex) upgradeWebsocket(w http.ResponseWriter, r *http.Request) {
	var err error
	g.websocketConn, err = g.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Could not upgrade websocket connection.")
		return
	}
	log.Printf("Connection established to gex editor")
}

func New(r *io.Reader, w *io.Writer) (gex *Gex) {

	//reader, writer := io.Pipe()

	html, err := Asset("app/dist/build.html")
	check(err)
	upgrader := &websocket.Upgrader{ReadBufferSize: 65535, WriteBufferSize: 65535}
	websocketMutex := &sync.Mutex{}
	g := &Gex{r, w, nil, websocketMutex, string(html), upgrader}

	http.HandleFunc("/", g.serveApp)
	http.HandleFunc("/ws", g.upgradeWebsocket)

	/*
		go func() {
			err = http.ListenAndServe("127.0.0.1:8081", nil)
			check(err)
		}()
	*/
	go func() {

		// TODO, currently modifying this to work from teh reader writer passed in
		// look at tcp dump
		ioutil.ReadAll(*r)
		time.Sleep(5 * time.Second)
		for {
			if g.websocketConn == nil {
				continue
			}
			messageType, moddedBytes, err := g.websocketConn.ReadMessage()
			log.Println(messageType)
			//if type == websocket.Message
			fmt.Println(moddedBytes)
			//websocketMutex.Unlock()
			if err != nil {
				log.Printf("[ERR] Failed to read from websocket: %v\n", err)
				continue
			}
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Second)
			if g.websocketConn == nil {
				continue
			}
			data := []byte("ABCD")
			if err := g.websocketConn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				log.Printf("[ERR] Failed to write to websocket: %v\n", err)
				//websocketMutex.Unlock()
				continue
			}
		}
	}()

	err = http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		panic(err)
	}
	return
}
