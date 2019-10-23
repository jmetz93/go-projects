package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// jacob := person{"Jacob", "Metzinger"} // one method of declaring structs, assigns values based on order
	jacob := person{firstName: "Jacob", lastName: "Metzinger"}
	fmt.Println(jacob)
}
