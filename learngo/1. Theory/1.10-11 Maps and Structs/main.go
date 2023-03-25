package main

import "fmt"

type person struct {
	name     string
	age      int
	favMusic []string
}

func main() {
	baskin := map[string]string{
		"name": "baskin",
		"age":  "24",
	}

	for key, value := range baskin {
		fmt.Println(key, value)
	}

	fmt.Println(baskin)

	favMusic := []string{"Clair de lune", "Confutatis"}
	baskin2 := person{name: "baskin", age: 24, favMusic: favMusic}
	fmt.Println(baskin2)
}
