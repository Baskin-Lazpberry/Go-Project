package main

import "fmt"

func main() {
	a := 2
	b := a
	c := &a
	fmt.Println(a, b, *c)

	a = 10
	fmt.Println(a, b, *c)

	*c = 8
	fmt.Println(a, b, *c)
}
