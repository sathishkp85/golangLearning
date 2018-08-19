package main

import "fmt"

func main() {
	a := 1
	b := true
	c := 1.1
	d := "hi there"

	fmt.Printf("%v, %v, %v, %v \n", a, b, c, d)
	fmt.Printf("%d, %t, %g, %q \n", a, b, c, d)
	fmt.Printf("%T, %T, %T, %T \n", a, b, c, d)
}
