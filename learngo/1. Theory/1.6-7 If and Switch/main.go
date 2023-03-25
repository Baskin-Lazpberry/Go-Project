package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	} else {
		return true
	}
}

func canIDrink2(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func amIUnemployed(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 70:
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(15))

	fmt.Println(canIDrink2(18))

	fmt.Println(amIUnemployed(24))
}
