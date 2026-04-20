// Scope.go
package main

import "fmt"

// Declared outside any function, visible across the whole package
var PackageVar = "I am at package scope"

func FunctionScopeDemo() {
	var functionVar = "I am inside FunctionScopeDemo"
	fmt.Println(functionVar) // accessible only inside this function
}
func BlockScopeDemo() {
	if true {
		blockVar := "I am inside an if-block"
		fmt.Println(blockVar) // accessible only inside this block
	}
	// fmt.Println(blockVar) // would be undefined here
}

func LoopScopeDemo() {
	for i := 0; i < 3; i++ {
		loopVar := i * 10
		fmt.Println("Loop iteration:", i, "Value:", loopVar)
	}
	// fmt.Println(loopVar) //  not accessible outside the loop
}

func ClosureDemo() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func ScopeDemo() {
	fmt.Println(PackageVar)
	FunctionScopeDemo()
	BlockScopeDemo()
	LoopScopeDemo()
	inc := ClosureDemo()
	fmt.Println(inc()) // prints 1
	fmt.Println(inc()) // prints 2
	inc2 := ClosureDemo()
	fmt.Println(inc2())
	fmt.Println(inc2())
}
