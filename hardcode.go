package convert

import (
	"fmt"
	"strings"
)

func convert(input := ">++++++++++++++++++++++++++++++++++++++++++++----.>..,,[[]]]]]?><<") {
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
