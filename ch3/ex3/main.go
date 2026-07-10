package main

import "fmt"

func main(){
	//Exercise:
	/*Write a program that defines a struct called Employee with three fields:
	firstName, lastName, and id. The first two fields are of type string, and the
	last field (id) is of type int. Create three instances of this struct using whatever
	values you’d like. Initialize the first one using the struct literal style without
	names, the second using the struct literal style with names, and the third with a
	var declaration. Use dot notation to populate the fields in the third struct. Print
	out all three structs.*/

	type employee struct{
		firstName string
		lastName string
		id int
	}

	noNames := employee{
		id:1,
	}

	dante := employee{
		firstName: "Dante",
		lastName: "None",
		id: 2,
	}

	var rbn employee
	rbn.firstName = "rbn"
	rbn.lastName = "avs"
	rbn.id = 3

	fmt.Println(noNames)
	fmt.Println(dante)
	fmt.Println(rbn)
}
