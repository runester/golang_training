package main

import "fmt"

func main() {
	inc := func() func() int {
		var x int
		return func() int {
			x++
			return x
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d := %d\n", i, inc())
	}
}
