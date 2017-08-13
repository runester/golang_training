package main

import (
	"fmt"
	"sort"
	"strings"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return strings.Compare(p[i], p[j]) == -1
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	list := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(list)
	sort.Sort(list)
	fmt.Println(list)
	sort.Sort(sort.Reverse(list))
	fmt.Println(list)
}
