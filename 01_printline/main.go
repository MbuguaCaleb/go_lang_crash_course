package main

import "fmt"

func main(){
	age :=35
	name :="Caleb"

	//print
	//print does not add a new line to the end of the String
	// fmt.Print("Hello, ")
	// fmt.Print("world")

	//To add a new Line we use a special escape character
	// fmt.Print("Hello, ")
	// fmt.Print("World \n")
	// fmt.Print("New Line \n")

    //Printlin
	//Automatically adds a new line
	fmt.Println("Hello World")
	fmt.Println("GoodBye Heroes!")


	//remember simpleGo rules..
	//eg strings must be in quotes
	//all varibles declared must be in use.
	fmt.Println("My age is", age,  "and my name is", name)


	//fornatted string
	//Pringf(formatted string)
	//%v--format specifier for  variables
	fmt.Printf("My age is %v and my name is %v \n",age,name)


	//%q-places quoyes around variables when we OUTPUT THEM
	//IDEALLU USED IN Strings
	fmt.Printf("My age is %q and my name is %q \n",age,name)
	fmt.Printf("Age is of type %T \n",age)
	fmt.Printf("You scored %0.1f points! \n", 225.55)


	//Sprintf (SAVES Formatted strings and returns them in form of a variable)
	var str= fmt.Sprintf("My age is %v and my name is %v \n",age,name)
	fmt.Println("The saved string is:",str)

}