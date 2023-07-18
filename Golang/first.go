package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// http.HandleFunc("/", htmlvsPlain)
	// fmt.Print("The server is running at port 8080.")
	// http.ListenAndServe(":8080", nil)

	server := http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}

	// server.ListenAndServe()

	var muxGuardianMode http.ServeMux
	server.Handler = &muxGuardianMode
	muxGuardianMode.HandleFunc("/", helloWorldGuardianMode)
	http.HandleFunc("/timeout", timeout)

	server.ListenAndServe()
}

func helloHandler(write http.ResponseWriter, read *http.Request) {

	switch read.URL.Path {
	case "/":
		fmt.Fprint(write, "Hello World!")
	case "/guardian":
		fmt.Fprint(write, "Parthiv")
	default:
		fmt.Fprint(write, "Fundamental Error!")
	}
	fmt.Printf("Handling function with %s request\n", read.Method)
}

func htmlvsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlvsPlain")
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout Attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did *not* timeout")
}

func helloWorldGuardianMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWorldGuardianMode")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello World!</h1>")
}
