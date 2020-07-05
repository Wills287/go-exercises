package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func main() {
	http.HandleFunc("/", Chain(
		Hello,
		Logging(),
		Auth(),
	))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintln(w, "Hello world!")
	if err != nil {
		log.Fatalln("Failed to print")
	}
}

func Chain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	for x := len(m) - 1; x >= 0; x-- {
		f = m[x](f)
	}
	return f
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("LOGGING START")
			defer log.Println("LOGGING END")

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("Failed to read request body", err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("Request received: %v %v %q\n", r.URL.Path, r.Header, body)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			start := time.Now()

			defer func() {
				log.Printf("Request handled: %v, duration: %v", r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}

func Auth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("AUTH START")
			defer log.Println("AUTH END")
			a := r.Header.Get("Authorization")
			if !strings.Contains(a, "Basic") {
				http.Error(w, fmt.Sprintf("Requires basic authentication"), http.StatusUnauthorized)
				return
			}
			if !strings.Contains(a, "1234") {
				http.Error(w, fmt.Sprintf("Failed authentication"), http.StatusUnauthorized)
				return
			}
			f(w, r)
		}
	}
}
