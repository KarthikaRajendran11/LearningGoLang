package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, `says "Hello World  !"`)
}

func main() {
	xi := []int{2, 4, 7, 9, 42}
	fmt.Println(xi)

	m := map[string]int{
		"Karthika": 31,
		"navin":    28,
	}
	fmt.Println(m)

	p1 := person{
		"Mr",
		"MoneyPenny",
	}
	p1.speak()
}
