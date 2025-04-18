package common

import (
	"fmt"
	"testing"
)

func TestStringTruncate(t *testing.T) {
	var inputs = []struct {
		function func(string, int, string) string
		funcName string
		input    string
		maxSize  int
		talis    string
		expected string
	}{
		{TruncateText, "TruncateText", "Hello world", 4, "...", "H..."},
		{TruncateText, "TruncateText", "Hello world", 6, "...", "Hel..."},
		{TruncateText, "TruncateText", "Hello", 100, "...", "Hello"},
		{TruncateTextBeginning, "TruncateTextBeginning", "Hello world", 4, "...", "...d"},
		{TruncateTextBeginning, "TruncateTextBeginning", "Hello world", 6, "...", "...rld"},
		{TruncateTextBeginning, "TruncateTextBeginning", "Hello", 100, "...", "Hello"},
		{TruncateMiddleText, "TruncateMiddleText", "Hello world", 5, "...", "H...d"},
		{TruncateMiddleText, "TruncateMiddleText", "Hello world", 7, "...", "He...ld"},
		{TruncateMiddleText, "TruncateMiddleText", "Hello", 100, "...", "Hello"},
	}

	for _, tt := range inputs {
		t.Run(fmt.Sprintf("Run %s on string %s to %d chars", tt.funcName, tt.input, tt.maxSize), func(t *testing.T) {
			result := tt.function(tt.input, tt.maxSize, tt.talis)
			expected := tt.expected
			if result != expected {
				t.Errorf("got \"%s\", expected \"%s\"", result, expected)
			}
		})
	}
}

func TestFilenameWithouText(t *testing.T) {
	var inputs = []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"hello.zip", "hello"},
		{"hello.tar.gz", "hello"},
		{".gitignore", ".gitignore"},
		{"", ""},
	}

	for _, tt := range inputs {
		t.Run(fmt.Sprintf("Remove extension from %s", tt.input), func(t *testing.T) {
			result := FileNameWithoutExtension(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestIsBufferPrintable(t *testing.T) {
	var inputs = []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"hello", true},
		{"abcdABCD0123~!@#$%^&*()_+-={}|:\"<>?,./;'[]", true},
		{"Horizontal Tab and NewLine\t\t\n\n", true},
		{"\xa0(NBSP)", true},
		{"\x0b(Vertical Tab)", true},
		{"\x0d(CR)", true},
		{"ASCII control characters : \x00(NULL)", false},
		{"\x05(ENQ)", false},
		{"\x0f(SI)", false},
		{"\x1b(ESC)", false},
		{"\x7f(DEL)", false},
	}
	for _, tt := range inputs {
		t.Run(fmt.Sprintf("Testing if buffer %q is printable", tt.input), func(t *testing.T) {
			result := IsBufferPrintable([]byte(tt.input))
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsExtensionExtractable(t *testing.T) {
	inputs := []struct {
		ext      string
		expected bool
	}{
		{".zip", true},
		{".rar", true},
		{".7z", true},
		{".tar.gz", true},
		{".tar.bz2", true},
		{".exe", false},
		{".txt", false},
		{".tar", true},
		{"", false},    // Empty string case
		{".ZIP", true}, // Case sensitivity check
		{".Zip", true}, // Case sensitivity check
		{".bz", true},
		{".gz", true},
		{".iso", true},
	}

	for _, tt := range inputs {
		t.Run(tt.ext, func(t *testing.T) {
			result := IsExtensionExtractable(tt.ext)
			if result != tt.expected {
				t.Errorf("IsExensionExtractable (%q) = %v; want %v", tt.ext, result, tt.expected)
			}
		})
	}
}

func TestMakePrintable(t *testing.T) {
	var inputs = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "hello"},
		{"abcdABCD0123~!@#$%^&*()_+-={}|:\"<>?,./;'[]", "abcdABCD0123~!@#$%^&*()_+-={}|:\"<>?,./;'[]"},
		{"Horizontal Tab and NewLine\t\t\n\n", "Horizontal Tab and NewLine\t\t\n\n"},
		{"(NBSP)\xa0\xa0\xa0\xa0;", "(NBSP)\xa0\xa0\xa0\xa0;"},
		{"\x0b(Vertical Tab)", "(Vertical Tab)"},
		{"\x0d(CR)", "(CR)"},
		{"ASCII control characters : \x00(NULL)", "ASCII control characters : (NULL)"},
		{"\x05(ENQ)", "(ENQ)"},
		{"\x0f(SI)", "(SI)"},
		{"\x1b(ESC)", "(ESC)"},
		{"\x7f(DEL)", "(DEL)"},
	}
	for _, tt := range inputs {
		t.Run(fmt.Sprintf("Make %q printable", tt.input), func(t *testing.T) {
			result := MakePrintable(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
