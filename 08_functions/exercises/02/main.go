package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	for i := 1; i <= 25; i++ {
		v, e := half(i)
		fmt.Printf("%d ? %d, %v\n", i, v, e)
	}
}
