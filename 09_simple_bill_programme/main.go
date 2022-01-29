package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func getInput (prompt string, r *bufio.Reader) (string,error){
  fmt.Print(prompt)
  input,err := r.ReadString('\n')
  return strings.TrimSpace(input),err

}

func createBill () bill {
	//creating a reader that will read from the standard input
	reader :=  bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ",reader)

	//calling method to create a new Bill
	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s - save bill, t - add tip): ",reader)

	switch opt {
	case "a":
		name, _ := getInput("Item Name:", reader)
		price, _ := getInput("Item Price:", reader)
		
		//can only parse numbers formatted as Strings
		p, err :=strconv.ParseFloat(price, 64)

		if err != nil{
			fmt.Println("The price must be a Number")
			promptOptions(b)
		}

		b.addItem(name,p)
		fmt.Println("Item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ",reader)
		t, err :=strconv.ParseFloat(tip, 64)

		if err != nil{
			fmt.Println("The tip must be a Number")
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("tip added - ",tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)	
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}

func main() {
	myBill :=createBill()
	promptOptions(myBill)	
}