package main

import "fmt"

func makeOddGenerator() func() uint {
	var i uint = 1
	nextValue := func() uint {
		var ret uint = i
		i += 2
		return ret
	}
	return nextValue
}

func main() {
	nextElem := makeOddGenerator()
	fmt.Println(nextElem())
	fmt.Println(nextElem())
	fmt.Println(nextElem())
	fmt.Println(nextElem())
}
