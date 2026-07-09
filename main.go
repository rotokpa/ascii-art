package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("invalid usage : <test>")
		return
	}

	input := strings.Join(os.Args[1:], " ")
	lines := strings.Split(input, "\n")

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		rendard := GenerateArt(line, banner)
		fmt.Println(rendard)
	}
}
