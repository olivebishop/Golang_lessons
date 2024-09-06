package main
import "fmt" 

func main ()  {
	//declaring and assigning
	var swahiliFood = "Biriani"

	fmt.Println(swahiliFood)

	//reassiging
	swahiliFood = "mahamri n mitahi"
	
	fmt.Println(swahiliFood)


	var otherFavSwahiliFood = "chapati n beans"
	fmt.Println(otherFavSwahiliFood)
    


	var myBestSwahiliSnack string
	var myPhoneNumber int
	var amICoolAsFuck bool

	myBestSwahiliSnack = "kashata"

	fmt.Println( myBestSwahiliSnack,myPhoneNumber,amICoolAsFuck )
	
	//using compound creations
	// var  myHeight, favShoes, favSport = 6, "yezzy", "buyern" 

	//using block creations
var(
	myHeight = 6
	favShoes = "yezzy"
	favSport = "buyern"
)
   myHeight,favShoes,favSport = 6, "yezzy","buyern"
    
   fmt.Print(myHeight,favShoes,favSport)

   //best practices and shortcut ofusing var we use :=
    favPet := "cat"
	fmt.Println(favPet)

	// declaring using compound variations

	myEx1,myEx2,myEx3 := "wanjiku", "aisha","shinyakanga"
	//assign  & reassign my ex
	myEx3 , myEx4 := "pam", "aoko"
    fmt.Println(myEx1 ,myEx2,myEx3 , myEx4)


	//constants 
	//(remember always constants dont change)

	const myName = "olive"
	fmt.Println(myName)

    }  