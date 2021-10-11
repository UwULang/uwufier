package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ">++++++++++++++++++++++++++++++++++++++++++++----.>..,,[[]]]]]?><<"
	
	replaces := map[string]string {
		"+": "👆",
		"-": "👇",
		">": "👉",
		"<": "👈",
		".": "🥺",
		",": "😳",
		"?": "🥴",
		"[": "😒",
		"]": "😡",
	}
	
	for key, value := range replaces {
		input = strings.Replace(input, key, value , -1)
	}
	
	
	fmt.Println(input)
}
