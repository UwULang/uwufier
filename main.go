package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var input string
	if (len(os.Args) < 2) {
		// get from cin if provides no user args
		fmt.Scanln(&input)
	} else {
		// otherwise try opening the first args
		data, error := os.ReadFile(os.Args[1])
		if error != nil {
			log.Fatal(error)
		}
		input = string(data)
	}
	
	
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
