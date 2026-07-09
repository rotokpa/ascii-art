package main

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	input = strings.ReplaceAll(input, "\n", "\\n")
	w := SplitInput(input)
	var result strings.Builder
	for i, word := range w {
		if word == "" {
			if i == len(w)-1 {
				continue
			}
			result.WriteByte('\n')
			continue
		}
		g := RenderLine(word, banner)
		for i := range 8 {
			result.WriteString(g[i])
			result.WriteRune('\n')
		}
	}
	return result.String()
}
