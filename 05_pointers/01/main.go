package main

import "fmt"

func main() {
	var a = 43
	var b = &a

	fmt.Printf("a := %v at %v\n", a, &a)
	fmt.Printf("b := %v at %v\n", b, *b)

}
