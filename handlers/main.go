package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	w.Write([]byte("The time is: " + tm))
}

func timeHandleFunc(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC3339Nano)
	w.Write([]byte("The time is " + tm))
}

func timeHandlerClosure(format string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)
}

func main() {
	log.Println("Handlers in Golang")
	port := flag.String("port", ":4000", "port for the server to listen to")
	mux := http.NewServeMux() // Same as DefaultServeMux (get's instantiated by default)
	// avoid using DefaultServeMux as it could pose a security risk because of global scope

	rh := http.RedirectHandler("https://example.org", http.StatusTemporaryRedirect)
	th := &timeHandler{format: time.RFC1123Z}

	th2 := http.HandlerFunc(timeHandleFunc) // mux.HandleFunc("/time", timeHandler)

	thC := timeHandlerClosure(time.RFC3339)

	mux.Handle("/foo", rh)
	mux.Handle("/tiempo", th)
	mux.Handle("/time", th2)
	mux.Handle("/time-new", thC)

	log.Printf("Listening ... on http://localhost%s", *port)
	http.ListenAndServe(*port, mux)
}
