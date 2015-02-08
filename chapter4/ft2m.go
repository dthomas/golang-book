package main

import "fmt"

func main()  {
  var feet float64
  fmt.Println("Converting feet to meters")
  fmt.Print("Enter value in feet: ")
  fmt.Scanf("%f", &feet)
  meters := feet * 0.3048
  fmt.Println(meters, "m = ", feet, "ft")
}
