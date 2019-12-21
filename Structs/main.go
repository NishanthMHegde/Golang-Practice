package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//declaring an empty person
	var p1 person
	fmt.Println(p1)
	p2 := person{firstName: "Nishanth",
		lastName: "Hegde",
		contact: contactInfo{
			email:   "nhegde12@gmail.com",
			zipCode: 560056,
		},
	}
	fmt.Println("Printing person")
	p2.print()
	fmt.Println("Printing updated person")
	//ideal way of assigning the address of a struct to create pointer
	newPerson := &p2
	newPerson.updateName("Nishu")
	p2.print()

	//pass by value can also be done in this way by directly passing the struct as receiver.
	// Go will create a pointer for the struct automatically
	// p2.updateName("Nishuuu")
	// p2.print()
}

func (p person) print() {

	fmt.Printf("%+v", p)
}

//By default Go uses pass by value. In order to avoid that and to actually update the person name, we should use pass by reference
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//Important note on pointers

//Slices can be updated by passing them directly in receiver functions without conveting them to memory addresses.
//Slice is actually array in the backend. Most of the list of items are slices.
//This is because Slices consists of 3 fields: Pointer to head of array, capacity and current length of slice.
//When a slice is created, the slice is stored in a memory address and its corresponding array is stored on a different memory address.
//When a slice is passed as receiver, a copy of the slice is passed. Even though a copy of the slice is passed, its pointer will still point
// to the same array.
