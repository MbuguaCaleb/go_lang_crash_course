package main

import "fmt"

func main() {
   
	//strings in Go must be decalred in double quotes
	var nameOne string = "mario"

	//Go can also inferTypes
	var nameTwo= "luigi"

	//3rd way to declare a variable
	var nameThree string

	fmt.Println(nameOne,nameTwo,nameThree)

	nameOne = "caleb"
	nameThree = "mbugua"

	fmt.Println(nameOne,nameTwo,nameThree)


	//shorthand way to declare variable
    //Note however, this shorthand syntax cannot be used outside the function
	nameFour := "Karanja"
	fmt.Println(nameFour)

	//Integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 50

	fmt.Println(ageOne,ageTwo,ageThree)


	//bits and memory
	//we can assign our variabe size
	var numOne int8 = 25
	var numTwo int8 = -128

	//we are not including minuses
	//allminuses are now going the the plus side
	var numThree uint8 = 25


	//Float
	//Number with decimal points
	//we must specify the bit size
	var scoreOne float32 = 25.98
	var scoreTwo float64 =34678678634786786.54

	//Floatt 64 is inferred automatically
	//It is prefreed since it has got a higher precision
	scoreThree := 1.5

	fmt.Println(scoreOne,scoreTwo,scoreThree)

}