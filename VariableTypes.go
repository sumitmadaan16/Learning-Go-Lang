package main

import "fmt"

func VariableDemo() {
	fmt.Println("Basic Types")
	var a int = 1
	var b int8 = 2  // -128 to 127
	var c int16 = 3 // 	-32,768 to 32,767
	var d int32 = 4 // -2,147,483,648 to 2,147,483,647
	var e int64 = 5 // 	-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
	fmt.Println("d =", d)
	fmt.Println("e =", e)

	var f float32 = 3.14
	var g float64 = 4.56
	fmt.Println("f =", f)
	fmt.Println("g =", g)

	var h bool = true
	fmt.Println("h =", h)

	fmt.Println("Composite Data Types")

	//Array of type int
	// you can also declare (if not want to initialize) as
	// var i [3]int
	i := [3]int{1, 2} // third element will be 0
	fmt.Println("Array h = ", i)

	//Slice

	//"initially in Go these start with0 capacity"
	//Subsequent growth: capacity increases exponentially, often doubling (not 1.5x like Java).
	j := []int{1, 2, 3}
	j = append(j, 4, 5, 6) // add using append method
	more := []int{7, 8}    // can also add 2 slices together
	j = append(j, more...)
	fmt.Println("Slice/Dynamic Array j = ", j)
	// we don't have any remove function we have to reallocate array again
	j = append(j[:2], j[3:]...) // removes value at index 2 i.e (3)
	fmt.Println("reallocated Array j = ", j)

	//Structs
	type k struct {
		id   int
		name string
	}
	student := k{id: 1, name: "xyz"}
	fmt.Printf("Struct  k = %v\n", student)

	//Maps
	l := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(l)
	l["b"] = 99 // Update key "b" to new value
	fmt.Println(l)

}
