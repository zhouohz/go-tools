package char

import "unicode"

// IsASCII checks if a character is an ASCII character (0-127).
func IsASCII(ch rune) bool {
	return ch < 128
}

// IsASCIIPrintable checks if a character is a printable ASCII character (32-126).
func IsASCIIPrintable(ch rune) bool {
	return ch >= 32 && ch < 127
}

// IsASCIICtrl checks if a character is an ASCII control character (0-31 and 127).
func IsASCIICtrl(ch rune) bool {
	return ch < 32 || ch == 127
}

// IsLetter checks if a character is a letter (uppercase or lowercase).
func IsLetter(ch rune) bool {
	return IsLetterUpper(ch) || IsLetterLower(ch)
}

// IsLetterUpper checks if a character is an uppercase letter (A-Z).
func IsLetterUpper(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// IsLetterLower checks if a character is a lowercase letter (a-z).
func IsLetterLower(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

// IsNumber checks if a character is a digit (0-9).
func IsNumber(ch rune) bool {
	return unicode.IsDigit(ch)
}

// IsHexChar checks if a character is a hexadecimal character (0-9, a-f, A-F).
func IsHexChar(ch rune) bool {
	return IsNumber(ch) || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

// IsLetterOrNumber checks if a character is a letter (A-Z, a-z) or a number (0-9).
func IsLetterOrNumber(ch rune) bool {
	return IsLetter(ch) || IsNumber(ch)
}

// IsBlankChar checks if a character is a whitespace character (space, tab, newline, etc.).
func IsBlankChar(ch rune) bool {
	return unicode.IsSpace(ch) || ch == '\ufeff' || ch == '\u202a' || ch == '\u0000' || ch == '\u3164' || ch == '\u2800' || ch == '\u180e'
}

// IsEmoji checks if a character is an emoji.
func IsEmoji(ch rune) bool {
	return (ch != 0x0) && (ch != 0x9) && (ch != 0xA) && (ch != 0xD) && ((ch >= 0x20 && ch <= 0xD7FF) || (ch >= 0xE000 && ch <= 0xFFFD) || (ch >= 0x100000 && ch <= 0x10FFFF))
}

// IsFileSeparator checks if a character is a file separator (slash or backslash).
func IsFileSeparator(ch rune) bool {
	return ch == '/' || ch == '\\'
}

// EqualsIgnoreCase checks if two characters are equal, optionally ignoring case.
func EqualsIgnoreCase(c1, c2 rune, caseInsensitive bool) bool {
	if caseInsensitive {
		return unicode.ToLower(c1) == unicode.ToLower(c2)
	}
	return c1 == c2
}
