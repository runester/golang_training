package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 43
	fmt.Printf("x := %v\n", x)
	zero(&x)
	fmt.Printf("x := %v\n", x)
}
