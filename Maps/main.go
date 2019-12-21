package main

import "fmt"

func main() {
	//Maps are like dictionaries in Python
	//In Maps, all keys need to be of same type and all values need to be of same type but keys and value types can de different

	//empty map declaration
	var map1 map[string]string
	fmt.Println(map1)

	//another way to declare empty map
	map2 := make(map[string]string)
	fmt.Println(map2)

	//declaring a full map
	map3 := map[string]string{
		"red":  "#80AF",
		"blue": "#009A",
	}

	fmt.Println(map3)

	fmt.Println("Let us add an item to map")
	map3["white"] = "#0000"
	fmt.Println(map3)
	fmt.Println("Let us delete an item in map")
	delete(map3, "blue")
	fmt.Println(map3)
	iterateMap(map3)
}

func iterateMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, " has a hex value of ", value)
	}
}

//Difference between Map and Struct:

// Map is used to store key value pairs which may not be related to each other. Structs store properties of same type of element

//Map is passed by reference whereas Struct is passed by value

//Struct can have different type of keys and values whereas all keys and all values should be of same type in Maps.
