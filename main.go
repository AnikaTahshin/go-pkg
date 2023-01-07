package main

import (
	"fmt"
	"unicode"
)

func UpperCase(input string) {

	for _, currentLetter := range input {

		if unicode.IsSpace(currentLetter) {
			fmt.Printf("%c", currentLetter)

		} else if unicode.IsLetter(currentLetter) {
			fmt.Printf("%c", currentLetter-32)
		} else {
			fmt.Printf("%c", currentLetter)
		}

	}

}
