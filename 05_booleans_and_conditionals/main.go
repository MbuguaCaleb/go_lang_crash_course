package main

import "fmt"

func main() {
	age := 25

	fmt.Println(age <=50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)


	//If else statements in Go Lang
	if age < 30 {
		fmt.Println("age is less than 30")
	}else if age < 40 {
		fmt.Println("Age is less than 40")
	}else{
		fmt.Println("Age is not less than 45")
	}

	//break and continue
	
	
}
