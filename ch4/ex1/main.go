package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Exercise:
	//Write a for loop that puts 100 random numbers between 0 and 100 into an int slice

	//Create the slice with a capacity of 100 numbers
	int_numbers := make([]int,0,100)

	//Populate the slice
	for i := 0; i<100; i++ {
		int_numbers = append(int_numbers, rand.Intn(100))
	}

	//print the numbers
	fmt.Println(int_numbers)
}
