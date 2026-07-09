package main 

import "fmt"


func ValidateInput(s string) (rune, error ) {
	for _, cha := range s {
		if cha < 32 || cha > 126 {
			return cha, fmt.Errorf("invalid :  %q", cha)
		}
	}
	return 0, nil
}