package main

import (
	"fmt"
	"net/http"
	"sync"
)

var waitRoom = sync.WaitGroup{}

func main() {

	http.HandleFunc("/", routines)
	fmt.Print("The server is running at port 8080.")
	http.ListenAndServe(":8080", nil)
}

func routines(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 10; i++ {

		waitRoom.Add(1)
		if i == 0 {
			go Add(w, r)
		} else if i == 1 {
			go Substract(w, r)
		} else if i == 2 {
			go Multiply(w, r)
		} else if i == 3 {
			go Divide(w, r)
		} else if i == 4 {
			go Modulo(w, r)
		} else if i == 5 {
			go squareRoot(w, r)
		} else if i == 6 {
			go powerOf(w, r)
		} else {
			go helloWorld(w, r)
		}
	}
	waitRoom.Wait()
}

func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Add")
	waitRoom.Done()
}

func Substract(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Substract")
	waitRoom.Done()
}

func Multiply(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Multiply")
	waitRoom.Done()
}

func Divide(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Divide")
	waitRoom.Done()
}

func Modulo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Modulo")
	waitRoom.Done()
}

func squareRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Square Root")
	waitRoom.Done()
}

func powerOf(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Power of")
	waitRoom.Done()
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World from go Routine")
	waitRoom.Done()
}
