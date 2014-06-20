package govalidator

import "testing"

func Test_Contains(t *testing.T) {
	if !Contains("abacada", "") {
		t.FailNow()
	}
	if Contains("abacada", "ritir") {
		t.FailNow()
	}
	if !Contains("abacada", "a") {
		t.FailNow()
	}
	if !Contains("abacada", "aca") {
		t.FailNow()
	}
}

func Test_Matches(t *testing.T) {
	if !Matches("123456789", "[0-9]+") {
		t.FailNow()
	}
	if Matches("abacaba", "cab$") {
		t.FailNow()
	}
	if !Matches("111222333", "((111|222|333)+)+") {
		t.FailNow()
	}
}

func Test_LeftTrim(t *testing.T) {
	if LeftTrim("  \r\n\tfoo  \r\n\t   ", "") != "foo  \r\n\t   " {
		t.FailNow()
	}
	if LeftTrim("010100201000", "01") != "201000" {
		t.FailNow()
	}
}

func Test_RightTrim(t *testing.T) {
	if RightTrim("  \r\n\tfoo  \r\n\t   ", "") != "  \r\n\tfoo" {
		t.FailNow()
	}
	if RightTrim("010100201000", "01") != "0101002" {
		t.FailNow()
	}
}

func Test_Trim(t *testing.T) {
	if Trim("  \r\n\tfoo  \r\n\t   ", "") != "foo" {
		t.FailNow()
	}
	if Trim("010100201000", "01") != "2" {
		t.FailNow()
	}
}

func Test_WhiteList(t *testing.T) {
	if WhiteList("abcdef", "abc") != "abc" {
		t.FailNow()
	}
	if WhiteList("aaaaaaaaaabbbbbbbbbb", "abc") != "aaaaaaaaaabbbbbbbbbb" {
		t.FailNow()
	}
	if WhiteList("a1b2c3", "abc") != "abc" {
		t.FailNow()
	}
	if WhiteList("   ", "abc") != "" {
		t.FailNow()
	}
}

func Test_BlackList(t *testing.T) {
	if BlackList("abcdef", "abc") != "def" {
		t.FailNow()
	}
	if BlackList("aaaaaaaaaabbbbbbbbbb", "abc") != "" {
		t.FailNow()
	}
	if BlackList("a1b2c3", "abc") != "123" {
		t.FailNow()
	}
	if BlackList("   ", "abc") != "   " {
		t.FailNow()
	}
}

func Test_StripLow(t *testing.T) {
	if StripLow("foo\x00", false) != "foo" {
		t.FailNow()
	}
	if StripLow("\x7Ffoo\x02", false) != "foo" {
		t.FailNow()
	}
	if StripLow("\x01\x09", false) != "" {
		t.FailNow()
	}
	if StripLow("foo\x0A\x0D", false) != "foo" {
		t.FailNow()
	}

	if StripLow("perch\u00e9", false) != "perch\u00e9" {
		t.FailNow()
	}
	if StripLow("\u20ac", false) != "\u20ac" {
		t.FailNow()
	}
	if StripLow("\u2206\x0A", false) != "\u2206" {
		t.FailNow()
	}

	if StripLow("foo\x0A\x0D", true) != "foo\x0A\x0D" {
		t.FailNow()
	}
	if StripLow("\x03foo\x0A\x0D", true) != "foo\x0A\x0D" {
		t.FailNow()
	}
}

func Test_ReplacePattern(t *testing.T) {
	if ReplacePattern("ab123ba", "[0-9]+", "aca") != "abacaba" {
		t.FailNow()
	}
	if ReplacePattern("abacaba", "[0-9]+", "aca") != "abacaba" {
		t.FailNow()
	}
	if ReplacePattern("httpftp://github.comio", "(ftp|io)", "") != "http://github.com" {
		t.FailNow()
	}
	if ReplacePattern("aaaaaaaaaa", "a", "") != "" {
		t.FailNow()
	}
}
