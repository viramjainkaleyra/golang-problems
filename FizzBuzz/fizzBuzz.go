package main

import (
	"fmt"
)

//Program to print Fizz if multiple of 3,Buzz if multiple of 5 and FizzBuzz if multiple of both 3 and 5
func main() {
	var n int
	fmt.Println("Enter value of n")
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
