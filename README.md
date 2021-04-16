# Gex
Gex is a small library that embeds a Vue.js web app that
provides hex and text editing as well as the ability to intercept, save, and inject
data into any stream. This library provides an easy to use and versatile interface for data manipulation. Gex is instantiated with a `PipeReader` and `PipeWriter`,
the consumer and producer for each pipe can be anything. For an example of how
I use Gex personally, see [my fork of tcpprox](https://github.com/jeffssh/tcpprox)
which includes a new gex shim on every connection. I've also included a supplemental test file at `cmd/test/main.go`.

## Screenshots
![Intercepting messages](https://github.com/jeffssh/gex/raw/main/imgs/intercept.png)
![Using repeater functionality](https://github.com/jeffssh/gex/raw/main/imgs/repeater.png)


## Example Use

```golang
// create two pipes, one for input and one for output
gexInR, gexInW := io.Pipe()
gexOutR, gexOutW := io.Pipe()
// create new gex with 2^16 buffer size
g, err = gex.New(gexInR, gexOutW, 2<<16)
if err != nil {
    // handle
    return
}
go func() {
    // serve the Vue.js app
    g.Serve()
}()
// io.Copy the left connection (client) onto the input pipe writer
go func() {
    io.Copy(gexInW, connL)
}()
// io.Copy the output pipe reader to the appropriate destination
```

