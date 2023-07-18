package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Person interface {
	getName() string
	getPhoneNumber() string
	getAdhaarNumber() string
}

type Student struct {
	name    string
	phone   string
	aadhaar string
}

// // CollegeStudent struct
// type CollegeStudent struct {
// 	Student
// 	collegeName string
// }

func main() {

	Student := createStudent()

	// Validate constraints
	if len(Student.getPhoneNumber()) != 10 {
		fmt.Println("Sorry, you must have a 10 digit phone number.")
		return
	}

	if len(Student.getAdhaarNumber()) != 12 {
		fmt.Println("Sorry, you must have a 12 digit aadhar number.")
		return
	}

	fmt.Printf("Welcome, %s!\n", Student.getName())
	fmt.Printf("Phone Number: %s\n", generateUUIDFromString(Student.getPhoneNumber()))
	fmt.Printf("Aadhar Number: %s\n", generateUUIDFromString(Student.getAdhaarNumber()))

}

func createStudent() Person {
	name := getUserInput("Please enter your name: ")

	phone := getUserInput("Please enter your phone number: ")
	phoneNumber := phone

	aadhar := getUserInput("Please enter your aadhaar number: ")
	aadharNumber := aadhar

	student := Student{
		name:    name,
		phone:   phoneNumber,
		aadhaar: aadharNumber,
	}

	fmt.Println(student)
	return student
}

func (s Student) getName() string {
	return s.name
}

func (s Student) getPhoneNumber() string {
	return s.phone
}

func (s Student) getAdhaarNumber() string {
	return s.aadhaar
}

func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func generateUUIDFromString(did string) string {
	hash := sha256.Sum256([]byte(did))
	uuid := hex.EncodeToString(hash[:])
	return uuid
}
