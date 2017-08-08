package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var rs []int
	for i := 0; i < 12; i++ {
		rs = append(rs, int(r.Float64()*-100)-1)
	}
	fmt.Println(rs)
	fmt.Printf("max: %d\n", max(rs...))
}

func max(ns ...int) int {
	var max int
	for i, n := range ns {
		if n > max || i == 0 {
			max = n
		}
	}
	return max
}
