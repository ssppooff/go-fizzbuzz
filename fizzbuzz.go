package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Print(i, ": ")
		fizzy := ""

		if i%3 == 0 {
			fizzy = fizzy + "Fizz"
		}

		if i%5 == 0 {
			fizzy = fizzy + "Buzz"
		}

		if fizzy == "" {
			fmt.Println(i)
		} else {
			fmt.Println(fizzy)
		}
	}
}
