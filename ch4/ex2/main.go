package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Exercise
	/*Loop over the slice you created in exercise 1. For each value in the slice, apply the
	following rules:
	a. If the value is divisible by 2, print “Two!”
	b. If the value is divisible by 3, print “Three!”
	c. IIf the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
	d. Otherwise, print “Never mind”.*/

	//EXERCISE 1
	int_numbers := make([]int,0,100)

	for i := 0; i<100; i++ {
		int_numbers = append(int_numbers, rand.Intn(100))
	}
	fmt.Println(int_numbers)

	//EXERCISE 2:

	for _,v := range int_numbers{
		switch{

		case v % 2 == 0 && v % 3 == 0:
			
			fmt.Println("Six - ",v)

		case v % 2 == 0:
			fmt.Println("Two - ",v)

		case v % 3 == 0:
			fmt.Println("Three - ",v)

		default:
			fmt.Println("Never mind - ",v)
		}
	}


}
