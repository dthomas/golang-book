package main

import "fmt"

func halfeven(x int) (int, bool) {
	half := x / 2
	even := x%2 == 0
	return half, even
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, " => ")
		fmt.Println(halfeven(i))
	}
}
