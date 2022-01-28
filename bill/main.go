package main

import "fmt"

//It is like classes

func main() {

	myBill := newBill("caleb's bill")

	fmt.Println(myBill.format())
}