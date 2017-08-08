package main

import "fmt"

func foo(n ...int) {
	if len(n) > 0 {
		fmt.Printf("%v\n", n)
	} else {
		fmt.Println("no arguments")
	}
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
