package main

import (
	"errors"
	"fmt"
	"strconv"
)

//program to demonstrate error handling
func main() {
	var n int
	fmt.Println("Enter value of n")
	fmt.Scanln(&n)
	if n%5 != 0 {
		errorMessage := "Error -> " + strconv.Itoa(n) + " is not a multiple of 5"
		fmt.Println(errors.New(errorMessage))
		return
	}
	fmt.Println("Success ->", n, "is a multiple of 5")
}
