package govalidator

import (
	"regexp"
	"strings"
)

// Contains check if the string contains the substring.
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}

// Matches check if string matches the pattern (pattern is regular expression)
// In case of error return false
func Matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// LeftTrim trim characters from the left-side of the input.
// If second argument is empty, it's will be remove leading spaces.
func LeftTrim(str, chars string) string {
	pattern := ""
	if chars == "" {
		pattern = "^\\s+"
	} else {
		pattern = "^[" + chars + "]+"
	}
	r, _ := regexp.Compile(pattern)
	return string(r.ReplaceAll([]byte(str), []byte("")))
}

// RightTrim trim characters from the right-side of the input.
// If second argument is empty, it's will be remove spaces.
func RightTrim(str, chars string) string {
	pattern := ""
	if chars == "" {
		pattern = "\\s+$"
	} else {
		pattern = "[" + chars + "]+$"
	}
	r, _ := regexp.Compile(pattern)
	return string(r.ReplaceAll([]byte(str), []byte("")))
}


// Trim trim characters from both sides of the input.
// If second argument is empty, it's will be remove spaces.
func Trim(str, chars string) string {
	return LeftTrim(RightTrim(str, chars), chars)
}

// WhiteList remove characters that do not appear in the whitelist.
func WhiteList(str, chars string) string {
	pattern := "[^" + chars + "]+"
	r, _ := regexp.Compile(pattern)
	return string(r.ReplaceAll([]byte(str), []byte("")))
}

// BlackList remove characters that appear in the blacklist.
func BlackList(str, chars string) string {
	pattern := "[" + chars + "]+"
	r, _ := regexp.Compile(pattern)
	return string(r.ReplaceAll([]byte(str), []byte("")))
}

// StripLow remove characters with a numerical value < 32 and 127, mostly control characters.
// If KeepNewLines is true, newline characters are preserved (\n and \r, hex 0xA and 0xD).
func StripLow(str string, KeepNewLines bool) string {
	chars := ""
	if KeepNewLines {
		chars = "\x00-\x09\x0B\x0C\x0E-\x1F\x7F"
	} else {
		chars = "\x00-\x1F\x7F"
	}
	return BlackList(str, chars)
}

// ReplacePattern replace regular expression pattern in string
func ReplacePattern(str, pattern, replace string) string {
	r, _ := regexp.Compile(pattern)
	return string(r.ReplaceAll([]byte(str), []byte(replace)))
}
