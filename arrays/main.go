package main

import (
	"fmt"
)

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr[1:3]) //we use for loop for this in java

	var intSlice []int32 = []int32{4, 5, 6}
	//fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	//fmt.Println(intSlice)
	//fmt.Printf("The length is %v and capacity is %v", len(intSlice), cap(intSlice)) //4 and cap 7

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) //appending two arrays
	fmt.Println(intSlice)

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

}
