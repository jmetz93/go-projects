package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// jacob := person{"Jacob", "Metzinger"} // one method of declaring structs, assigns values based on order
	// jacob := person{firstName: "Jacob", lastName: "Metzinger"}
	// var jacob person // creates struct with empty property values, in case of person struct each prop will be empty string
	// fmt.Println(jacob)
	// jacob.firstName = "Jacob"
	// jacob.lastName = "Metzinger"
	// fmt.Printf("%+v", jacob) // logs struct with property names and values
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// jimPointer := &jim // creating pointer to memory address of struct
	jim.updateFirstName("Jacob")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// *person ensures to use the memory address of the pointer
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	// (*) says to grab actual struct from memory
	(*pointerToPerson).firstName = newFirstName
}
