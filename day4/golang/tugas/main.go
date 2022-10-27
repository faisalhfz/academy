package main

import "fmt"

func shape(n int) {
	if n%2 == 0 {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")
		}
	} else {
		for i := 1; i <= n; i++ {
			for j := 1; j <= i; j++ {
				fmt.Print("*")
			}
			fmt.Print("\n")
		}
	}
}

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	shape(n)
}
