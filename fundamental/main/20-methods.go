package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

}

/**
 * Go supports methods defined on struct types.
 * This area method has a receiver type of *rect.
 * Methods can be defined for either pointer or value receiver types. Here’s an example of a value receiver.
 * Here we call the 2 methods defined for our struct.
 * Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct.
 * Next we’ll look at Go’s mechanism for grouping and naming related sets of methods: interfaces.
 */
