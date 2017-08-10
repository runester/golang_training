package main

import (
	"fmt"
	"sort"
)

func keys(m map[string]int) []string {
	ks := make([]string, len(m))
	var i int
	for k := range m {
		ks[i] = k
		i++
	}
	sort.Strings(ks)
	return ks
}

func main() {
	fruit := map[string]int{
		"Banana":    12,
		"Cherry":    8,
		"Apple":     10,
		"Blueberry": 14,
	}
	fmt.Println(fruit)
	fmt.Println(keys(fruit))
}
