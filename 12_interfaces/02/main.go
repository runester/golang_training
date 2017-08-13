package main

import (
	"fmt"
	"reflect"
)

type car struct {
	name string
}

type boat struct {
	name string
}

type person struct {
	name string
}

type list interface{}

func listing(l ...list) {
	for _, v := range l {
		fmt.Println(v)
	}
}

func main() {
	cars := []car{car{"Grey Ford"}, car{"White CRV"}, car{"Blue Truck"}}
	listing(cars)

	boats := []boat{boat{"Sunsetter"}, boat{"Watercraft"}, boat{"Skidoo"}}
	listing(boats)

	persons := []person{person{"Abel"}, person{"Baker"}, person{"Charlie"}}
	listing(persons)

	ints := []int{2, 5, 9}
	listing(ints)

	floats := []float32{1.11, 2.22, 3.33}
	listing(floats)

	var masterlist []list
	for _, v := range cars {
		masterlist = append(masterlist, v)
	}
	for _, v := range boats {
		masterlist = append(masterlist, v)
	}
	for _, v := range persons {
		masterlist = append(masterlist, v)
	}
	for _, v := range ints {
		masterlist = append(masterlist, v)
	}
	for _, v := range floats {
		masterlist = append(masterlist, v)
	}
	listing(masterlist)

	fmt.Println("cars are ", reflect.TypeOf(cars))
	fmt.Println("boats are ", reflect.TypeOf(boats))
	fmt.Println("persons are ", reflect.TypeOf(persons))
	fmt.Println("ints are ", reflect.TypeOf(ints))
	fmt.Println("floats are ", reflect.TypeOf(floats))
	fmt.Println("masterlist are ", reflect.TypeOf(masterlist))
}
