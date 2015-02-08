package main

import "fmt"

func main()  {
  var meters float64
  fmt.Println("Converting meters to feet")
  fmt.Print("Enter value in meters: ")
  fmt.Scanf("%f", &meters)
  feet := meters / 0.3048
  fmt.Println(meters, "m = ", feet, "ft")
}
