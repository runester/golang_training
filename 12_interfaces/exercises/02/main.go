package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(list)
	sort.Ints(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println(list)
}
