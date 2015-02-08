package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(i, "! = ", factorial(uint(i)))
		defer fmt.Println(i, "! = ", factorial(uint(i)))
	}
}
