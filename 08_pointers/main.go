package main

import (
	"fmt"
)

//a function that takes in a pointer to a string
//then references it
func updateName(x *string) {
	*x="Mbugua"
}

func main()  {

	name:= "caleb"

	//fmt.Println("memory address of name is: ", &name)

	//stroring a Poiter to the Memory Location
	//Moemiry address
	m := &name

	//we can dereference pointer to get  the Value it points to get the value
   //fmt.Println("Value at Memory address:", *m)

   //When i update my Pointer it is going to Update the variable it is pointing to
	updateName(m)

	fmt.Println(name)
	
}