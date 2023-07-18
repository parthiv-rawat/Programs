package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for idx, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", idx, day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco
		}

		if rogueValue == 5 {
			rogueValue++
			continue
			// break
		}

		fmt.Println("Value is", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.in")
}
