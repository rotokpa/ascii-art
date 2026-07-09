package main

import (
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// Load a real banner for testing
	banner, err := LoadBanner("standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}

	t.Run("empty input produces empty string", func(t *testing.T) {
		result := GenerateArt("", banner)
		if result != "" {
			t.Errorf("Expected empty string, got %q", result)
		}
	})

	t.Run("single newline produces one blank line", func(t *testing.T) {
		result := GenerateArt("\n", banner)
		if result != "\n" {
			t.Errorf("Expected %q, got %q", "\n", result)
		}
	})

	t.Run("single word produces 8 lines each ending with newline", func(t *testing.T) {
		result := GenerateArt("Hi", banner)
		lines := splitByNewline(result)
		// splitByNewline on "line1\nline2\n...\nline8\n" gives 8 non-empty + 1 empty from trailing \n
		// so we expect 9 elements with last being ""
		if len(lines) != 9 {
			t.Errorf("Expected 9 elements (8 art lines + trailing empty), got %d", len(lines))
		}
		// Each of the 8 art lines must be non-empty
		for i := 0; i < 8; i++ {
			if lines[i] == "" {
				t.Errorf("Line %d should not be empty for input 'Hi'", i)
			}
		}
		// Last element after trailing \n split should be empty
		if lines[8] != "" {
			t.Errorf("Expected trailing empty element after split, got %q", lines[8])
		}
	})

	t.Run("A newline B produces 8 lines for A then 8 lines for B no extra blank", func(t *testing.T) {
		result := GenerateArt("A\nB", banner)
		lines := splitByNewline(result)
		// 8 lines for A + 8 lines for B + trailing empty = 17
		if len(lines) != 17 {
			t.Errorf("Expected 17 elements (16 art lines + trailing empty), got %d: %v", len(lines), lines)
		}
		for i := 0; i < 16; i++ {
			if lines[i] == "" {
				t.Errorf("Line %d should not be empty for 'A\\nB'", i)
			}
		}
	})

	t.Run("A double-newline B produces 8 lines for A + 1 blank line + 8 lines for B", func(t *testing.T) {
		result := GenerateArt("A\n\nB", banner)
		lines := splitByNewline(result)
		// 8 lines for A + 1 blank line + 8 lines for B + trailing empty = 18
		if len(lines) != 18 {
			t.Errorf("Expected 18 elements (8 + blank + 8 + trailing), got %d", len(lines))
		}
		// lines[0..7] = A art
		for i := 0; i < 8; i++ {
			if lines[i] == "" {
				t.Errorf("Line %d (A art) should not be empty", i)
			}
		}
		// lines[8] = blank line
		if lines[8] != "" {
			t.Errorf("Line 8 should be blank (empty segment), got %q", lines[8])
		}
		// lines[9..16] = B art
		for i := 9; i < 17; i++ {
			if lines[i] == "" {
				t.Errorf("Line %d (B art) should not be empty", i)
			}
		}
		// lines[17] = trailing empty from final \n
		if lines[17] != "" {
			t.Errorf("Expected trailing empty element, got %q", lines[17])
		}
	})

	t.Run("empty segment produces exactly one blank line not eight", func(t *testing.T) {
		result := GenerateArt("A\n\nB", banner)
		lines := splitByNewline(result)
		// Confirm line 8 is the blank line (not 8 blank lines)
		blankCount := 0
		for i := 8; i < len(lines)-1; i++ {
			if lines[i] == "" {
				blankCount++
			} else {
				break
			}
		}
		if blankCount != 1 {
			t.Errorf("Expected exactly 1 blank line for empty segment, got %d", blankCount)
		}
	})

	t.Run("multiple consecutive newlines produce multiple blank lines", func(t *testing.T) {
		result := GenerateArt("\n\n\n", banner)
		// "\n\n\n" splits into ["", "", "", ""] → 3 empty segments + trailing
		// Each empty segment → 1 blank line → "\n\n\n"
		if result != "\n\n\n" {
			t.Errorf("Expected %q, got %q", "\n\n\n", result)
		}
	})

	t.Run("single character produces 8 art lines", func(t *testing.T) {
		result := GenerateArt("A", banner)
		lines := splitByNewline(result)
		if len(lines) != 9 {
			t.Errorf("Expected 9 elements for single char, got %d", len(lines))
		}
	})

	t.Run("output ends with newline for non-empty input", func(t *testing.T) {
		result := GenerateArt("Hello", banner)
		if len(result) == 0 || result[len(result)-1] != '\n' {
			t.Errorf("Expected output to end with newline, got %q", result)
		}
	})

	t.Run("newline only input does not produce art lines", func(t *testing.T) {
		result := GenerateArt("\n", banner)
		lines := splitByNewline(result)
		// "\n" → one blank line → split gives ["", ""]
		if len(lines) != 2 {
			t.Errorf("Expected 2 elements for single newline input, got %d", len(lines))
		}
		if lines[0] != "" {
			t.Errorf("Expected blank line, got %q", lines[0])
		}
	})
}

// splitByNewline splits a string by "\n" (preserving empty strings from consecutive newlines)
func splitByNewline(s string) []string {
	if s == "" {
		return []string{}
	}
	result := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	result = append(result, s[start:])
	return result
}