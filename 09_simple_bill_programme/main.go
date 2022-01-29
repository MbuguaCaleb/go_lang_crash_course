package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Println(opt)


}

func main() {
	myBill :=createBill()
	promptOptions(myBill)


	
}