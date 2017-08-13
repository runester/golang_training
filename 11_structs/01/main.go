package main

import "fmt"

type auto struct {
	make  string
	model string
	year  int
}

type person struct {
	first string
	last  string
	age   int
	car   auto
}

func main() {
	p1 := person{"Stephen", "Jarjoura", 47, auto{"Ford", "Escape", 2006}}
	p2 := person{"Lisa", "Jarjoura", 48, auto{"Honda", "CRV", 2014}}
	var p3 person
	p3.first = "Beau"
	p3.last = "Jarjoura"
	p3.age = 2
	fmt.Println(p1, p2, p3)
	fmt.Println(p1.first, p1.last, p1.age, p1.car.make, p1.car.model, p1.car.year)
}
