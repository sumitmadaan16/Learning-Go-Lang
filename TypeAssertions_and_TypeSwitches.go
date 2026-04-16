package main

import "fmt"

func TypeAssertionsDemo() {
	// we create an empty interface, and it is different from conventional interfaces as those have methods which can be implemented, but in this data varible can hold any type of data. gives us flexibility to fetch data when the type is unknown
	var data interface{} = "hello this is an example of type assertion"
	//  here we are telling " i think data has a type string, go can you check" if it is string then go return ok = true, else it sets ok as false

	// If you omit , ok (like str := data.(string)), then Go assumes you’re 100% sure. If you’re wrong, the program panics.
	str, ok := data.(string)
	if ok { // here it returns true as it actually is true
		fmt.Println("String Value", str)
	} else {
		fmt.Println("Not a String")
	}
	// but here it returns false, as we say data has type as int, but when go checks it it is a string
	num, ok := data.(int)
	if ok {
		fmt.Println("Numerical value: ", num)
	} else {
		fmt.Println("Not a numerical value")
	}
	// modifying the data
	data = 42
	// Type switch example
	switch v := data.(type) { //in these cases this switch comes handy as it automatically checks among multiple data types
	case string:
		fmt.Println("Switch says: String Value:", v)
	case int:
		fmt.Println("Switch says: Numerical Value:", v)
	case float64:
		fmt.Println("Switch says: Float Value:", v)
	default:
		fmt.Println("Switch says: Unknown type")
	}

}
