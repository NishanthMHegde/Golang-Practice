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

	fmt.Printf("%+v", p1)
	p2 := person{firstName: "Nishanth",
		lastName: "Hegde",
		contact: contactInfo{
			email:   "nhegde12@gmail.com",
			zipCode: 560056,
		},
	}
	fmt.Println(p2)
}
