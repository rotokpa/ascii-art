package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune] []string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ")
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("error")
	}
	if len(data) < 95*9 {
		return nil, fmt.Errorf("error")
	}

	lines := strings.Split(string(data), "\n")
	listmap := make(map[rune][]string)
	for i := range 95 {
		cha := rune(i + 32)
		start := i*9 + 1
		window := lines[start: start+8]
		lines := make([]string, 8)
		copy(lines, window)
		listmap[cha] = lines
	}
	return listmap, nil 
}

