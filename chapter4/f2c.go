package main

import "fmt"

func main()  {
  var fahrenheit float64
  fmt.Println("Convert Fahrenheit to Celcious")
  fmt.Print("Enter temprature in Fahrenheit: ")
  fmt.Scanf("%f", &fahrenheit)
  celcius := (fahrenheit - 32) * (5.0/9.0)
  fmt.Println(fahrenheit, " in Celcius is: ", celcius)
}
