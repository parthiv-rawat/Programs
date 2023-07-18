package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for:", languages["JS"])
	fmt.Println("RB shorts for:", languages["RB"])
	fmt.Println("PY shorts for:", languages["PY"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// loops are intersting in golang

	for key, value := range languages {
		fmt.Printf("For key: %v, values is %v\n", key, value)
	}
}
