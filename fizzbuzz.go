package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, ": ")
		fizzy := false

		if i%3 == 0 {
			fmt.Print("Fizz")
			fizzy = true
		}

		if i%5 == 0 {
			fmt.Print("Buzz")
			fizzy = true
		}

		if !fizzy {
			fmt.Print(i)
		}

		fmt.Println()
	}
}
