package main

import (
	"fmt"

	"github.com/Baskin-Lazpberry/learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{}

	err := dictionary.Add("bonjour", "greeting")
	if err != nil {
		fmt.Println(err)
	}
	printWord(dictionary, "bonjour")

	err = dictionary.Add("bonjour", "greeting")
	if err != nil {
		fmt.Println(err)
	}

	dictionary.Update("bonjour", "good morning")
	printWord(dictionary, "bonjour")

	dictionary.Delete("bonjour")
	printWord(dictionary, "bonjour")
}

func printWord(dict dict.Dictionary, word string) {
	definition, err := dict.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
