package main

import (
	"fmt"
	"github.com/runester/shapes"
)

func info(s shapes.Shape) {
	fmt.Printf("%T %v\n", s, s)
	fmt.Println(s, s.Area())
}

func main() {
	sq := shapes.Square{Side: 10}
	info(sq)
	c1 := shapes.Circle{Radius: 3}
	info(c1)
}
