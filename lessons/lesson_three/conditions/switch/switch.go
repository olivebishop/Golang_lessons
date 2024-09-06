package main

import "fmt"

func main(){
	//change animal name and see the results in the terminal
	animal := "cat"

	switch animal{
	    case "cat":
			fmt.Println("meow")
		case "dog":
			fmt.Println("barks")
		case "horse":
			fmt.Println("neigh")
		case "snake":
			fmt.Println("hisss")

		default:
			fmt.Println("Unware of the animal dude!")


	}
}