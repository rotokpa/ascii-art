package main

import (
	"strings"
)

func SplitInput(text string) []string {
	word := strings.Split(text, "\\n")

	return word
}
