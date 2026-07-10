package main

import "fmt"

func main(){
	//Exercise:
	/*Write a program that defines a string variable called message with the value "Hi
	👩 and 👨" and prints the fourth rune in it as a character, not a number*/

	message := "Hi 👩 and 👨"
	var runeMessage []rune = []rune(message)

	fmt.Println(string(runeMessage[3]))


}
