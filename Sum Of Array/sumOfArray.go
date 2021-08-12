package main

import (
	"fmt"
)

//program to print sum of array elements
func main() {
	var n int
	fmt.Println("Enter number of elements")
	fmt.Scanln(&n)
	nums := make([]int, n)
	fmt.Println("Enter",n,"elements")
	for i := 0; i < n; i++ {
		fmt.Scanln(&nums[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	fmt.Println("Sum =", sum)
}
