package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 43
	m["k2"] = 44
	m["k3"] = 0
	for k, v := range m {
		fmt.Println(k, v)
	}
	s := []m
}
