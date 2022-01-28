package main 

import "fmt"

//a function is made up of what it taks an thereturn type it has
func updateName(x string) string{
	x="Caleb"
	return x 
}

func updateMenu (y map[string]float64){
y["coffee"] = 2.99
}

func main (){
	
	//Group a types ->strings,ints,bools,floats,arrays,structs
	name := "Joshua"

	//everytime we pass a variable into a function
	//Go creates a copy of the Variable
	//A copy is created and stored in its own memory location
	//this is why i must do reassignment
	name = updateName(name)

	fmt.Println(name)



	//Group B Types -->slices,maps,functions
	menu := map[string]float64{
		"pie":5.95,
		"ice cream":390,
	}

	updateMenu(menu)
	fmt.Println(menu)

}