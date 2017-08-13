package main

import "fmt"

type auto struct {
	make  string
	model string
	year  int
}

func (a auto) String() string {
	return fmt.Sprintf("%d %s %s", a.year, a.make, a.model)
}

type person struct {
	first string
	last  string
	age   int
	car   auto
}

func (p person) String() string {
	return fmt.Sprintf("%s %s (%d) drives a %s", p.first, p.last, p.age, p.car)
}

func main() {
	p1 := person{"Stephen", "Jarjoura", 47, auto{"Ford", "Escape", 2006}}
	p2 := person{"Lisa", "Jarjoura", 48, auto{"Honda", "CRV", 2014}}
	var p3 person
	p3.first = "Beau"
	p3.last = "Jarjoura"
	p3.age = 2
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
