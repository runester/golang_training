package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		v, e := half(i)
		fmt.Printf("%d ? %d, %v\n", i, v, e)
	}
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
