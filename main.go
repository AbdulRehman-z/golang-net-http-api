package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/message" {
		w.Write([]byte("Invaid roure"))
		return
	}

	if r.Method != "GET" {
		w.Write([]byte("Invalid method"))
		return
	}

	w.Write([]byte("message"))
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/form" {
		w.Write([]byte("invalid route"))
		return
	}

	if r.Method != "GET" {
		w.Write([]byte("Invalid method"))
		return
	}

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")

	fmt.Fprintf(w, "name: %s", name)

	w.Write([]byte("Abdul Rehman"))

}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/username", formHandler)
	fmt.Printf("Listening on port 8000")
	http.ListenAndServe(":8000", nil)
}
