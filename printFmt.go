package main

import "fmt"

func main() {
	fmt.Println("Hello go")
	fmt.Println(123)
	fmt.Printf("%x \n", 255)
	fmt.Printf("%#X \n", 255)
	fmt.Printf("%#x \n", 255)
	fmt.Printf("%b \n", 255)
	fmt.Printf("%#o \n", 255)

	for i := 2949; i < 2999; i++ {
		fmt.Printf("%d, %b, %o, %#x, %#q \n", i, i, i, i, i)
	}
}
