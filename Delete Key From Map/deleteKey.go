package main

import (
	"fmt"
)

//program to delete a key from map
func main(){
	myMap := make(map[int]int)
	myMap[0]=0
	myMap[1]=1
	myMap[2]=2
	fmt.Println("Map before deleting key ->",myMap)
	delete(myMap,1)
	fmt.Println("Map after deleting key ->",myMap)
}