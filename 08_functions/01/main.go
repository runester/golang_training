package main

import "fmt"

func main() {
	set := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	fmt.Println(filter(set, func(n int) bool {
		return n%2 == 0
	}))
}

func filter(num []int, cb func(int) bool) []int {
	var match []int
	for _, n := range num {
		if cb(n) {
			match = append(match, n)
		}
	}
	return match
}
