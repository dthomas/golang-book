package main

import "fmt"

func main() {
  x := []int{ 83, 34, 38, 28, 57, 9, 29, 57, 28, 20, 27, 12 }
  smallest := x[0]

  for _, value := range x {
    if value < smallest {
      smallest = value
    }
  }

  fmt.Println("Smallest value among ", x, "is ", smallest)
}
