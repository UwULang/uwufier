package convert

import (
	"strings"
)

func Convert(input string) (output string) {
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
		"\n": "",
		" ": "",
	}
	
	for key, value := range replaces {
		input = strings.Replace(input, key, value , -1)
	}
	
	return input
}
