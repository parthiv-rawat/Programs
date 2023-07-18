package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang ; No super or parent

	parthiv := User{"parthiv", "parthiv@gmail.com", true, 18}
	fmt.Println(parthiv)
	fmt.Printf("Parthiv details are: %+v\n", parthiv)
	fmt.Printf("Name is %v and email is %vn", parthiv.Name, parthiv.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
