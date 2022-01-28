package main

import "fmt"

func main() {
	// age := 25

	// fmt.Println(age <=50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)


	//If else statements in Go Lang
	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// }else if age < 40 {
	// 	fmt.Println("Age is less than 40")
	// }else{
	// 	fmt.Println("Age is not less than 45")
	// }

	//break and continue
	//Continue is used to Go back top of Loop in the current Iteration
	//Break is used to exit out of a Loop Completely
	names := []string{"caleb", "mbugua","Karanja","Kerei"}

	for index, value := range names {

		if index == 1 {
			fmt.Println("Contunuing at Pos", index)
			continue
		}

		if index > 2 {
			fmt.Println("Breaking at position",index)
			break
		}

		fmt.Printf("The Value at pos %v is %v \n", index,value)


	}

	
}
