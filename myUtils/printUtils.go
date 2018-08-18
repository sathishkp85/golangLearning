package myUtils

import "fmt"

func PrintAscii() {
	for i := 0; i < 128; i++ {
		fmt.Printf("%#x %c \n", i, i)
	}
}
