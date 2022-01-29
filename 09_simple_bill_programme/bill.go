package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bills function
//this function returns a bill type
func newBill(name string) bill {
	//intial bill Values
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

//receiver function
//format the Bill
func (b bill) format() string {

	fs := "Bill breakdown :\n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v  \n", k+":",v)
		total += v
	}

	//addTip
	fs += fmt.Sprintf("%-25v ...$%v\n","tip:",b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:",total+b.tip)

	return fs

}

//update  tip
//structs are also treated as classA therfore we must pass the pointers
//Go automatically does referencing for pointers in structs
func (b *bill) updateTip(tip float64){
	b.tip = tip
}

//add an item to the bill
func (b *bill) addItem(name string, price float64){
	b.items[name]=price
}

//save bill
//save data to a file
func (b *bill) save(){

	//converting my Bill Map into a Slice
	data := []byte(b.format())

	//save bill into a file
	err := os.WriteFile("bills/"+b.name+".txt",data,0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")

}