package main

import "fmt"

func PointerDemo() {
	var p *int
	x := 10
	p = &x
	fmt.Printf("Address of x : %d\n", p)
	fmt.Printf("Value of x : %d\n", *p)

	// pointer arithmetic
	*p = pointerIncrement(p)
	fmt.Printf("Value of x : %d\n", *p)
	*p = pointerIncrement(&x)
	fmt.Printf("Value of x : %d\n", *p)

	// Pointer to pointer
	pp := &p
	fmt.Printf("Address of p: %d\n", pp)
	fmt.Printf("Address of x : %d\n", *pp)
	fmt.Printf("Value of x : %d\n", **pp)
}

func pointerIncrement(n *int) int {
	return *n + 1
}
