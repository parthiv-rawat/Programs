package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	// no inheritance in golang ; No super or parent

	parthiv := User{"parthiv", "parthiv@gmail.com", true, 18}
	fmt.Println(parthiv)
	fmt.Printf("Parthiv details are: %+v\n", parthiv)
	fmt.Printf("Name is %v and email is %v\n", parthiv.Name, parthiv.Email)

	parthiv.GetStatus()
	parthiv.NewMail()
	fmt.Printf("Name is %v and email is %v\n", parthiv.Name, parthiv.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
