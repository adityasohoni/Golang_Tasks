package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	for i := 1; i < 10000; i++ {
		if i%3 == 0 && i%5 == 0 && i%13 == 0 && i%7 == 0 {
			fmt.Print(i)
		}
	}

}
