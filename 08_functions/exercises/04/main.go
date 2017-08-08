package main

import "fmt"

func main() {
	if (true && false) || (false && true) || !(false && false) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
