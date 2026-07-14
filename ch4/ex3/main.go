package main

import "fmt"

func main(){
	//Exercise
	/*Start a new program. In main, declare an int variable called total. Write a for
	loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
	The body of the for loop should be as follows:
	total := total + i
	fmt.Println(total)*/

	var total int

	for i:= 0; i<10; i++{
		total := total + i //should use a '=' instead of ':='
		fmt.Println(total)
	}

	//After the for loop, print out the value of total. What is printed out? What is the likely bug in this code?
	fmt.Println(total) // <- This will print 0 if we didn't assign a value or the initial value we assigned
	//The for is shadowing the total variable, because of ':=', a new variable total is created.
	
	//Correct loop
	fmt.Println()
	for i:= 0; i<10; i++{
		total = total + i //Deleted ':' to use the same variable
		fmt.Println(total)
	}

	fmt.Println("After the loop: ",total)
}
