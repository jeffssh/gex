package gex

import (
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
}

func (g *Gex) checkFatal(err error) {
	if err != nil {
		log.Printf("[GEX] Fatal error: %v\n", err)
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
		}
		if g.listener != nil {
			g.listener.Close()
		}
	})
}

func (g *Gex) serveApp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string(g.html))
}

func (g *Gex) Serve() {
	url := "http://" + g.listener.Addr().String()
	log.Printf("[GEX] Serving editor at: %s\n", url)
	err := http.Serve(g.listener, g.router)
	g.checkFatal(err)
}

func (g *Gex) upgradeWebsocket(w http.ResponseWriter, r *http.Request) {
	var err error
	g.websocketConn, err = g.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[GEX] Could not upgrade websocket connection.")
		return
	}
	go func() {
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
	}()
	log.Printf("[GEX] Editor connected over ws://")
}

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

	g = &Gex{r, w, nil, &once, string(html), upgrader, listener, router}
	s := router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		// deny any host besides 127.0.0.1:*
		// to prevent DNS rebinding and cross-site websocket hijacking
		return r.Host == g.listener.Addr().String()
	}).Subrouter()
	s.HandleFunc("/", g.serveApp).Methods(http.MethodGet)
	s.HandleFunc("/ws", g.upgradeWebsocket).Methods(http.MethodGet)

	go func() {
		buff := make([]byte, 65536)
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
			g.checkFatal(err)
		}
	}()

	return
}
