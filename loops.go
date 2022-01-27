package main

import "fmt"

func main() {

	//all loops in go start with for
	//while fo loop
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value of x is",x)
	// 	x ++
	// }

	//for loop-traditional
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of i is", i)

	// }

	 names := []string{"caleb","mbugua","Karanja","Kerei","Ngugi"}

	// for i:=0 ; i<len(names); i++{
	// 	fmt.Println("value of is is", names[i])
	// }

	//for in Loop
	// for index, value := range names {
	// 	fmt.Printf("The value and index %v is %v \n", index,value)
	// }

	for _,value := range names {

		fmt.Printf("The value is %v \n", value)
		value = "NEW NAMES"
	}

	fmt.Println(names)
	

}

