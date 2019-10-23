package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// jacob := person{"Jacob", "Metzinger"} // one method of declaring structs, assigns values based on order
	jacob := person{firstName: "Jacob", lastName: "Metzinger"}
	// var jacob person // creates struct with empty property values, in case of person struct each prop will be empty string
	// fmt.Println(jacob)
	fmt.Printf("%+v", jacob) // logs struct with property names and values

}
