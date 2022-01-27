package main

import "fmt"

func main() {

	//Array rules in Go
	//It always must have a specific length
	//cannot be chnages after it has been declared
	//var ages [3]int = [3]int{20, 25, 30}

	//arrays in Go are immutables---sets in Python

	//shorthand array sytax
	//Notice the right hand assignment//
	//Very crucial and must be there
	var ages = [3]int{20, 25, 30}

	//Shorthand array assignment
	names := [4]string{"Caleb","Mbugua","Karanja","Kerei"}

	//we can also change the value of an array index
	names[1] = "wanjiru"

	//len prints length of arrays
	fmt.Println(ages, len(ages))
	fmt.Println(names,len(names))


	//slices---are like arrays like we know from normal programming Languages
	//We do not begin by putting in a number to check the Lenght of the Array
	var scores = []int{100,50,60}
	scores[2] = 25

	//append to slice
	//Append by itself just like the spread Operator returns a new Slice
	//Therefore i must reassign my new variable since it does not do that for me automatically
	//Appending is one thing you can be able to do in slices but not in arrays
	scores = append(scores, 85)

	fmt.Println(scores,len(scores))


	//slice ranges
	//Just like in Python
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree :=names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)


	rangeOne = append(rangeOne, "Wamboi")

   	fmt.Println(rangeOne)


}