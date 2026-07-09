package main

import (
	"strings"

	"testing"
)

// helper: loads standard banner or fails the test immediately.

func loadStandard(t *testing.T) map[rune][]string {

	t.Helper()

	banner, err := LoadBanner("standard.txt")

	if err != nil {

		t.Fatalf("could not load standard.txt: %v", err)

	}

	return banner

}

// TestRenderLine_ReturnsEightLines checks that any non-empty input always

// yields a slice of exactly 8 strings.

func TestRenderLine_ReturnsEightLines(t *testing.T) {

	banner := loadStandard(t)

	result := RenderLine("Hello", banner)

	if len(result) != 8 {

		t.Errorf("expected 8 lines, got %d", len(result))

	}

}

// TestRenderLine_EmptyStringReturnsEightEmptyLines checks that an empty input

// still produces 8 strings, all empty — not a nil slice.

func TestRenderLine_EmptyStringReturnsEightEmptyLines(t *testing.T) {

	banner := loadStandard(t)

	result := RenderLine("", banner)

	if len(result) != 8 {

		t.Errorf("expected 8 lines for empty input, got %d", len(result))

	}

	for i, line := range result {

		if line != "" {

			t.Errorf("line %d: expected empty string, got %q", i, line)

		}

	}

}

// TestRenderLine_SingleSpace checks that a single space character produces 8

// lines whose content matches the space character's art in the banner.

func TestRenderLine_SingleSpace(t *testing.T) {

	banner := loadStandard(t)

	result := RenderLine(" ", banner)

	if len(result) != 8 {

		t.Fatalf("expected 8 lines for space, got %d", len(result))

	}

	spaceArt := banner[' ']

	for i := range result {

		if result[i] != spaceArt[i] {

			t.Errorf("line %d: expected %q, got %q", i, spaceArt[i], result[i])

		}

	}

}

// TestRenderLine_TwoCharsConcatenated checks that rendering "AB" on each row

// equals the row for 'A' concatenated with the row for 'B'.

func TestRenderLine_TwoCharsConcatenated(t *testing.T) {

	banner := loadStandard(t)

	result := RenderLine("AB", banner)

	if len(result) != 8 {

		t.Fatalf("expected 8 lines, got %d", len(result))

	}

	for i := 0; i < 8; i++ {

		want := banner['A'][i] + banner['B'][i]

		if result[i] != want {

			t.Errorf("row %d: expected %q, got %q", i, want, result[i])

		}

	}

}

// TestRenderLine_NoTrailingNewline checks that none of the 8 returned strings

// ends with a newline character. Printing is the caller's job.

func TestRenderLine_NoTrailingNewline(t *testing.T) {

	banner := loadStandard(t)

	result := RenderLine("Hello", banner)

	for i, line := range result {

		if strings.HasSuffix(line, "\n") {

			t.Errorf("line %d ends with \\n — RenderLine must not add newlines", i)

		}

	}

}

// TestRenderLine_SingleCharMatchesBanner checks that rendering a single

// character produces exactly what is stored for that character in the banner.

func TestRenderLine_SingleCharMatchesBanner(t *testing.T) {

	banner := loadStandard(t)

	chars := []rune{'H', 'e', 'l', 'o', '1', '!', ' '}

	for _, r := range chars {

		result := RenderLine(string(r), banner)

		if len(result) != 8 {

			t.Errorf("char %q: expected 8 lines, got %d", r, len(result))

			continue

		}

		for i := 0; i < 8; i++ {

			if result[i] != banner[r][i] {

				t.Errorf("char %q row %d: expected %q, got %q", r, i, banner[r][i], result[i])

			}

		}

	}

}

// TestRenderLine_LongerWordWidth checks that the width of each row grows with

// more characters. "AB" rows must be wider than "A" rows.

func TestRenderLine_LongerWordWidth(t *testing.T) {

	banner := loadStandard(t)

	single := RenderLine("A", banner)

	double := RenderLine("AB", banner)

	for i := 0; i < 8; i++ {

		if len(double[i]) < len(single[i]) {

			t.Errorf("row %d: 'AB' should be wider than 'A', got %d <= %d",

				i, len(double[i]), len(single[i]))

		}

	}

}

// TestRenderLine_SpecialCharacters checks that special printable characters

// (punctuation, brackets, symbols) are handled without panic or empty output.

func TestRenderLine_SpecialCharacters(t *testing.T) {

	banner := loadStandard(t)

	specials := []string{"!", "#", "$", "%", "&", "()", "{}"}

	for _, s := range specials {

		result := RenderLine(s, banner)

		if len(result) != 8 {

			t.Errorf("input %q: expected 8 lines, got %d", s, len(result))

		}

	}

}
