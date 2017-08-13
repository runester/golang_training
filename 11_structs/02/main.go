package main

import "fmt"

type auto struct {
	make  string
	model string
	year  int
}

func (a auto) drive() {
	fmt.Printf("%s goes \"Vroom, Vroom!\"\n", a.make)
}

func (a auto) pretty() {
	fmt.Printf("%d %s %s\n", a.year, a.make, a.model)
}

func (a *auto) inc() {
	a.year++
}

func main() {
	var c1 auto
	c1.make = "Ford"
	c1.model = "Escape"
	c1.year = 2006
	c1.drive()
	c1.pretty()
	c1.inc()
	c1.pretty()
}
