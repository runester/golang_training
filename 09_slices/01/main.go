package main

import "fmt"

func main() {
	ds := make([]rune, 0, 128)
	for i := 0; i <= 58; i++ {
		ds = append(ds, rune('A'+i))
		fmt.Printf("%d - %s - (%d,%d) - %v\n", i, string(ds[i]), len(ds), cap(ds), string(ds))
	}
	fmt.Println(string(ds[5:12]))
}
