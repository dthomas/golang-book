package main

import "fmt"

func fibonacci(x uint) uint {
	switch {
	case x == 0:
		return 0
	case x == 1:
		return 1
	default:
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(fibonacci(uint(i)))
	}
}
