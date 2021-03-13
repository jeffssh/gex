//go:generate go run -tags generate gen.go
package main

import (
	"log"
	"os"
	"os/signal"
	"runtime"

	"github.com/zserge/lorca"
)

var B64html string

func main() {
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	ui, err := lorca.New("", "", 480, 320, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("start", func() {
		log.Println("UI is ready")
	})

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	ui.Load("data:text/html;base64," + B64html)
	//ui.Load("data:text/html," + editor)

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	ui.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
	/*
		ui, err := lorca.New("", "", 480, 320)
		if err != nil {
			log.Fatal(err)
		}
		defer ui.Close()

		// Data model: number of ticks
		ticks := uint32(0)
		// Channel to connect UI events with the background ticking goroutine
		togglec := make(chan bool)
		// Bind Go functions to JS
		ui.Bind("toggle", func() { togglec <- true })
		ui.Bind("reset", func() {
			atomic.StoreUint32(&ticks, 0)
			ui.Eval(`document.querySelector('.timer').innerText = '0'`)
		})

		// Load HTML after Go functions are bound to JS
		ui.Load("data:text/html," + url.PathEscape(`


		<html>
			<body>
				<script>alert(1)</script>
				<!-- toggle() and reset() are Go functions wrapped into JS -->
				<div class="timer" onclick="toggle()"></div>
				<button onclick="reset()">Reset</button>
			</body>
		</html>
		`))

		// Start ticker goroutine
		go func() {
			t := time.NewTicker(100 * time.Millisecond)
			for {
				select {
				case <-t.C: // Every 100ms increate number of ticks and update UI
					ui.Eval(fmt.Sprintf(`document.querySelector('.timer').innerText = 0.1*%d`,
						atomic.AddUint32(&ticks, 1)))
				case <-togglec: // If paused - wait for another toggle event to unpause
					<-togglec
				}
			}
		}()
		<-ui.Done()
	*/
}

const editor string = `<!-- this wonderful page was found here: https://github.com/xem/hex -->
<body onload='
// Reset the textarea value
m.value="00";
// Init the top cell content
for(i=0;i<16;i++)
  t.innerHTML+=(0+i.toString(16)).slice(-2)+" ";
'>
<!-- TRUDY SPECIFIC CODE ADDED FOR THIS PROJECT -->
<h1> ~ Trudy Intercept ~ </h1>
<script>
    var url = window.location.href
    var arr = url.split("/");
    var ws_url = "ws://" + arr[2] + "/ws"
    var socket = new WebSocket(ws_url)
    socket.onmessage = function (event) {
	document.getElementById('m').value = event.data
	document.getElementById('m').oninput()
	document.getElementById('send').disabled = false
    }
    var sender = function() {
        socket.send(document.getElementById('m').value)
	document.getElementById('send').disabled = true
        document.getElementById('m').value = "00"
        document.getElementById('m').oninput()
    }
</script>
<button onclick="sender()" id='send' disabled=true>send</button>
<!-- END TRUDY SPECIFIC CODE -->
</body>
<table border><td><pre><td id=t><tr><td id=l width=80>00000000<td><textarea spellcheck=false id=m oninput='
// On input, store the length of clean hex before the textarea caret in b
b=value
.substr(0,selectionStart)
.replace(/[^0-9A-F]/ig,"")
.replace(/(..)/g,"$1 ")
.length;
// Clean the textarea value
value=value
.replace(/[^0-9A-F]/ig,"")
.replace(/(..)/g,"$1 ")
.replace(/ $/,"")
.toUpperCase();
// Set the height of the textarea according to its length
style.height=(1.5+value.length/47)+"em";
// Reset h
h="";
// Loop on textarea lines
for(i=0;i<value.length/48;i++)
  
  // Add line number to h
  h+=(1E7+(16*i).toString(16)).slice(-8)+" ";
// Write h on the left column
l.innerHTML=h;
// Reset h
h="";
// Loop on the hex values
for(i=0;i<value.length;i+=3)
  
  // Convert them in numbers
  c=parseInt(value.substr(i,2),16),
  
  // Convert in chars (if the charCode is in [64-126] (maybe more later)) or ".".
  h=63<c&&127>c?h+String.fromCharCode(c):h+".";
  
// Write h in the right column (with line breaks every 16 chars)
r.innerHTML=h.replace(/(.{16})/g,"$1 ");
// If the caret position is after a space or a line break, place it at the previous index so we can use backspace to erase hex code
if(value[b]==" ")
  b--;
// Put the textarea caret at the right place
setSelectionRange(b,b)'
cols=48></textarea><td width=160 id=r>.</td>
</table>
<style>
*{margin:0;padding:0;vertical-align:top;font:1em/1em courier}
#m{height:1.5em;resize:none;overflow:hidden}
#t{padding:0 2px}
#w{position:absolute;opacity:.001}
</style>
`
