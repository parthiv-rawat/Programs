package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// commonChan := make(chan string, 2)
	chanAtoB := make(chan string, 5)
	chanBtoC := make(chan string, 5)
	chanCtoD := make(chan string, 5)

	var wg sync.WaitGroup
	wg.Add(15)
	for i := 0; i < 5; i++ {
		chanAtoB <- "b"
		go func(chanAtoB chan string) {
			chanBtoC <- "c"
			fmt.Printf("This function convert 'a' to '%v'\n", <-chanAtoB)
			wg.Done()

		}(chanAtoB)

		go func(chanBtoC chan string) {
			chanCtoD <- "d"
			fmt.Printf("This function convert 'b' to '%v'\n", <-chanBtoC)
			wg.Done()
		}(chanBtoC)

		go func(chanCtoD chan string) {
			fmt.Printf("This function convert 'c' to '%v'\n", <-chanCtoD)
			wg.Done()
		}(chanCtoD)

		time.Sleep(time.Second)
		fmt.Printf("The character at the end is d\n")
		fmt.Println("---------------")
	}

	// }

	wg.Wait()
}
