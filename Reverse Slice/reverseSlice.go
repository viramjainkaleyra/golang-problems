package main

import "fmt"

//program to reverse a slice order
func main() {
	nums := [...]int{1,2,3,4,5}
	fmt.Println("Before : ", nums)
	reversed := []int{}
	for i := range nums {
		reversed = append(reversed, len(nums)-i)
	}
	fmt.Println("After : ", reversed)
}