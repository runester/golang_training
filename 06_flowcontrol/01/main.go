package main

import "fmt"

func main() {
	i := 0
	for i < 100 {
		fmt.Printf("i := %d\n", i)
		i += 2
	}
	fmt.Printf("final value of i := %d\n", i)
}
