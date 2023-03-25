package main

import "fmt"

func main() {
	names := [5]string{"Baskin", "Wiki", "Lazpberry"}
	names[3] = "KTA"
	names[4] = "Jeong"
	fmt.Println(names)

	scores := []int{55, 76}
	scores = append(scores, 40)
	scores = append(scores, 98)
	scores = append(scores, 61)
	fmt.Println(scores)
}
