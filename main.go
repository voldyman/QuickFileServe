package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port      = flag.Int("port", 8000, "Port to run the server on")
	is_global = flag.Bool("global", true, "Should the file server be accessable from outside")
	folder    = flag.String("fpath", ".", "Which folder to serve")
)

func init() {
	flag.Parse()
}
func main() {
	addr := "%s:%d"

	if *is_global {
		addr = fmt.Sprintf(addr, "0.0.0.0", *port)
	} else {
		addr = fmt.Sprintf(addr, "127.0.0.1", *port)
	}

	fmt.Println("Starting server on address", addr)
	fmt.Println("Serving folder", *folder)

	http.Handle("/", newServeAndLog(*folder))

	http.ListenAndServe(addr, nil)
}

type ServeAndLog struct {
	handler http.Handler
}

func newServeAndLog(loc string) *ServeAndLog {
	return &ServeAndLog{
		handler: http.FileServer(http.Dir(loc)),
	}
}

func (m *ServeAndLog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	m.handler.ServeHTTP(w, r)
}
