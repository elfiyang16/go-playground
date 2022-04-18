package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo // if the name and property are the same, can ignore
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	var alex person // init alex without assigning it
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex) // {firstName: lastName:}%

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94500,
		}, // have to have the comma
	}
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//can either call with a person type or adr of person
func (pointerToPerson *person) updateName(newFirstName string) { // path by value
	(*pointerToPerson).firstName = newFirstName

}
