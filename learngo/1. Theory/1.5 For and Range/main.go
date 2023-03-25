package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func superAdd2(numbers ...int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func main() {
	result := superAdd(8, 9, 3, 0)
	fmt.Println(result)

	result2 := superAdd2(8, 9, 3, 1)
	fmt.Println(result2)
}
