package main

import "fmt"

func max(args ...int) int {
	maximum := args[0]
	for _, value := range args {
		if value > maximum {
			maximum = value
		}
	}
	return maximum
}

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(max(xs...))
}
