package main

import (
	"fmt"
	"sort"
)

func main() {

	//Strings Packaage
	//greeting := "Hello there friends"

	//Go Secret
	//When we change something we must do reassignment for the change to take Place
	// fmt.Println(strings.Contains(greeting,"hello"))
	// fmt.Println(strings.ReplaceAll(greeting,"Hello","hi"))
	// fmt.Println(strings.ToUpper(greeting))

	//find index of a certain element in a String
	// fmt.Println(strings.Index(greeting, "th"))

	// fmt.Println(strings.Split(greeting, ""))

	//Sorting the Slice
	ages := []int{45, 20, 35, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages ,90)

	fmt.Println(index)

	names := []string{"yes","no","caleb","peris"}

	sort.Strings(names)
	fmt.Println(names)

	//search through strings
	fmt.Println(sort.SearchStrings(names,"no"))


}