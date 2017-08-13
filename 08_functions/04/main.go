package main

import (
	"fmt"
	"strings"
)

func addStrings(s ...string) string {
	return strings.Join(s, " ")
}

func addInts(i ...int) int {
	var sum int
	for _, v := range i {
		sum += v
	}
	return sum
}

func main() {
	s1 := "Happy"
	s2 := "Days"
	fmt.Println(s1, s2, addStrings(s1, s2, "are", "here", "again"))
	i1 := 8
	i2 := 12
	fmt.Println(i1, i2, addInts(i1, i2, 2, 4, 7, 11))
}
