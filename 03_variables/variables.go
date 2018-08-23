package main

import "fmt"

func main() {
	a := 1
	b := true
	c := 1.1
	d := "hi there"

	var (
		v1 = 1
		v2 = "@"
		v3 = false
		v4 = 1.9
	)

	const MSG = "Welocme to the world of Go."
	fmt.Println(MSG)
	var testStr string
	fmt.Scanf("%s", &testStr)
	//testStr = "Hi There"
	fmt.Printf("%v, %v, %v, %v \n", a, b, c, d)
	fmt.Printf("%d, %t, %g, %q, %s \n", a, b, c, d, testStr)
	fmt.Printf("%T, %T, %T, %T \n", v1, v2, v3, v4)
}
