package main

import "fmt"

func main() {
	//arithmetic operators
	// + - * / %
 
	var num1  = 4
	var num2  = 2
	
	var sum = num1 + num2
	var diffence = num2-num1
	var quotience = num1/num2
	var product = num1*num2
	 var reminder = num1 % 2

	fmt.Println(sum, diffence, quotience,product,reminder)

    //relational operators
	var result = num1 < num2
	fmt.Println(result)

	//logical opearators
	const name = "aisher"
	var age = 24
	//resticted party invite with only  name with aisher or age above 24
	var invitedToParty = name == "aisher" && age > 24
    
	fmt.Println(invitedToParty)


	//not restricted party inite

	var invitedToPartyToday = name == "aisher"|| age > 24
    fmt.Println(invitedToPartyToday)

	

}
