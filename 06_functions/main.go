package main

import "fmt"

//creating our own funtions
func sayGreeting(n string){
	fmt.Printf("Good Morning %v \n",  n)
}

func sayBye(n string){
	fmt.Printf("Good Bye %v \n",  n)
}

func main(){
	sayGreeting("Caleb")
	sayGreeting("Harun")
	sayBye("David")
}