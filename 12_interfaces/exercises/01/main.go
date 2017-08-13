package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(list)
	sort.Strings(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(sort.StringSlice(list)))
	fmt.Println(list)
}
