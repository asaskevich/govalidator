package govalidator

import "testing"

func Test_Contains(t *testing.T) {
	tests_1 := []string{"abacada", "abacada", "abacada", "abacada"}
	tests_2 := []string{"", "ritir", "a", "aca"}
	expected := []bool{true, false, true, true}
	for i := 0; i < len(tests_1); i ++ {
		result := Contains(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_Matches(t *testing.T) {
	tests_1 := []string{"123456789", "abacada", "111222333", "abacaba"}
	tests_2 := []string{"[0-9]+", "cab$", "((111|222|333)+)+", "((123+]"}
	expected := []bool{true, false, true, false}
	for i := 0; i < len(tests_1); i ++ {
		result := Matches(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_LeftTrim(t *testing.T) {
	tests_1 := []string{"  \r\n\tfoo  \r\n\t   ", "010100201000"}
	tests_2 := []string{"", "01"}
	expected := []string{"foo  \r\n\t   ", "201000"}
	for i := 0; i < len(tests_1); i ++ {
		result := LeftTrim(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
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
