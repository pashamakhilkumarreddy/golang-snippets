package main

import (
	"flag"
	"io"
	"log"
	"mime"
	"net/http"
	"os"

	"github.com/goji/httpauth"
	"github.com/gorilla/handlers"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("Esta Bien!"))
}

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}
			if mt != "application/json" {
				http.Error(w, "Content-Type header must be of application/json", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func newLoggingHandler(dst io.Writer) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return handlers.LoggingHandler(dst, h)
	}
}

func main() {
	log.Println("Middleware in Golang")
	port := flag.String("port", ":4000", "port number for the server to listen on")
	flag.Parse()
	mux := http.NewServeMux()

	logFile, err := os.OpenFile("server.log",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatal(err)
	}
	loggingHandler := newLoggingHandler(logFile)
	authHandler := httpauth.SimpleBasicAuth("john", "p@ssw0rd")

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))
	mux.Handle("/json", enforceJSONHandler(finalHandler))
	mux.Handle("/auth", authHandler(finalHandler))
	mux.Handle("/logs", loggingHandler(finalHandler))
	log.Printf("The server is listening on ... http://localhost%s", *port)
	serverErr := http.ListenAndServe(*port, mux)
	log.Fatal(serverErr)
}
