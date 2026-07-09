package main

func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := range 8 {
		for _, cha := range text {
		result[i] += banner[cha][i]
		}
	}
	return result
}
