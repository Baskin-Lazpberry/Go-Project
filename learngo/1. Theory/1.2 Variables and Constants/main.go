package main

import "fmt"

func main() {
	const constName string = "Baskin"
	fmt.Println(constName)

	var varName string = "Raspberry"
	varName = "Lazpberry"
	fmt.Println(varName)

	//Only inside of function!
	name := "Wiki"
	fmt.Println(name)
}
