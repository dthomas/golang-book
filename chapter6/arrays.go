package main

import "fmt"

func main() {
  var total float64 = 0
  x := []float64{ 64, 36, 96, 95, 100, }
  for _, value := range x {
    total += value
  }

  fmt.Println(total/float64(len(x)))
}
