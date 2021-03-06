package gex

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Gex struct {
	r             *io.PipeReader
	w             *io.PipeWriter
	websocketConn *websocket.Conn
	once          *sync.Once
	html          string
	upgrader      *websocket.Upgrader
	listener      net.Listener
	router        *mux.Router
	buffSize      int
}

func (g *Gex) checkFatal(err error) {
	if err != nil {
		g.Close()
	}
}

func (g *Gex) Close() {
	g.once.Do(func() {
		if g.r != nil {
			g.r.Close()
		}
		if g.w != nil {
			g.w.Close()
		}
		if g.websocketConn != nil {
			g.websocketConn.Close()
			g.websocketConn = nil
		}
		if g.listener != nil {
			g.listener.Close()
		}
		fmt.Printf("[GEX] Shutting down gex\n")
	})
	// should i return an error here?
}

func (g *Gex) serveApp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string(g.html))
}

func (g *Gex) Serve() (err error) {
	url := "http://" + g.listener.Addr().String()
	log.Printf("[GEX] Serving editor at: %s\n", url)
	err = http.Serve(g.listener, g.router)
	g.checkFatal(err)
	return
}

func (g *Gex) pipeWsToWriter() {
	for {
		if g.websocketConn == nil {
			break
		}
		_, data, err := g.websocketConn.ReadMessage()
		if err != nil {
			break
		}
		_, err = g.w.Write(data)
		g.checkFatal(err)
	}
}

func (g *Gex) pipeReaderToWs() {
	buff := make([]byte, g.buffSize)
	for {
		bytesRead, err := g.r.Read(buff)
		g.checkFatal(err)
		data := buff[:bytesRead]
		if g.websocketConn != nil {
			err := g.websocketConn.WriteMessage(websocket.BinaryMessage, data)
			if err == nil {
				continue
			} else {
				// ws disconnected
				g.websocketConn = nil
			}
		}
		_, err = g.w.Write(data)
		if err != nil {
			g.Close()
			return
		}
	}
}

func (g *Gex) upgradeWebsocket(w http.ResponseWriter, r *http.Request) {
	var err error
	g.websocketConn, err = g.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[GEX] Could not upgrade websocket connection.")
		return
	}
	go g.pipeWsToWriter()
	log.Printf("[GEX] Editor connected over ws://")
}

/*
func (g *Gex) UpdatePipes(r *io.PipeReader, w *io.PipeWriter) {
	g.ClosePipes()
	g.r = r
	g.w = w
	if w != nil && r != nil {
		go g.pipeReaderToWs()
	}
}
*/

func New(r *io.PipeReader, w *io.PipeWriter, buffSize int) (g *Gex, err error) {
	var once sync.Once
	html, err := Asset("app/dist/build.html")
	if err != nil {
		log.Printf("[GEX] Fatal error: %v\n", err)
		return
	}
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Printf("[GEX] Fatal error: %v\n", err)
		return
	}
	upgrader := &websocket.Upgrader{ReadBufferSize: buffSize, WriteBufferSize: buffSize}
	router := mux.NewRouter()

	g = &Gex{r, w, nil, &once, string(html), upgrader, listener, router, buffSize}
	s := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		// deny any host besides 127.0.0.1:*
		// to prevent DNS rebinding and cross-site websocket hijacking
		return r.Host == g.listener.Addr().String() && r.Header.Get("origin") == ""
	}).Subrouter()
	s.HandleFunc("/", g.serveApp).Methods(http.MethodGet)
	s.HandleFunc("/ws", g.upgradeWebsocket).Methods(http.MethodGet)

	if g.w != nil && g.r != nil {
		go g.pipeReaderToWs()
	}
	return
}
