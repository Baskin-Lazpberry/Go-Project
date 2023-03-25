package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) (int, int) {
	return a * b, a + b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lenAndUpper2(name string) (length int, upper string) {
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func lenAndUpper3(name string) (length int, upper string) {
	defer fmt.Println("Did you call me?")
	length = len(name)
	upper = strings.ToUpper(name)
	return
}

func main() {
	result, _ := multiply(2, 2)
	fmt.Println(result)

	totalLength, upperName := lenAndUpper("Baskin")
	fmt.Println(totalLength, upperName)

	repeatMe("Un", "Deux", "Trois")

	totalLength2, upperName2 := lenAndUpper2("Lazpberry")
	fmt.Println(totalLength2, upperName2)

	totalLength3, upperName3 := lenAndUpper3("Wiki")
	fmt.Println(totalLength3, upperName3)
}
