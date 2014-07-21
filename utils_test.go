package govalidator

import "testing"

func Test_Contains(t *testing.T) {
	tests_1 := []string{"abacada", "abacada", "abacada", "abacada"}
	tests_2 := []string{"", "ritir", "a", "aca"}
	expected := []bool{true, false, true, true}
	for i := 0; i < len(tests_1); i++ {
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
	for i := 0; i < len(tests_1); i++ {
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
	for i := 0; i < len(tests_1); i++ {
		result := LeftTrim(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_RightTrim(t *testing.T) {
	tests_1 := []string{"  \r\n\tfoo  \r\n\t   ", "010100201000"}
	tests_2 := []string{"", "01"}
	expected := []string{"  \r\n\tfoo", "0101002"}
	for i := 0; i < len(tests_1); i++ {
		result := RightTrim(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_Trim(t *testing.T) {
	tests_1 := []string{"  \r\n\tfoo  \r\n\t   ", "010100201000", "1234567890987654321"}
	tests_2 := []string{"", "01", "1-8"}
	expected := []string{"foo", "2", "909"}
	for i := 0; i < len(tests_1); i++ {
		result := Trim(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

// This small example illustrate how to work with Trim function.
func ExampleTrim() {
	// Remove from left and right spaces and "\r", "\n", "\t" characters
	println(Trim("   \r\r\ntext\r   \t\n", "") == "text")
	// Remove from left and right characters that are between "1" and "8".
	// "1-8" is like full list "12345678".
	println(Trim("1234567890987654321", "1-8") == "909")
}

func Test_WhiteList(t *testing.T) {
	tests_1 := []string{"abcdef", "aaaaaaaaaabbbbbbbbbb", "a1b2c3", "   ", "a3a43a5a4a3a2a23a4a5a4a3a4"}
	tests_2 := []string{"abc", "abc", "abc", "abc", "a-z"}
	expected := []string{"abc", "aaaaaaaaaabbbbbbbbbb", "abc", "", "aaaaaaaaaaaa"}
	for i := 0; i < len(tests_1); i++ {
		result := WhiteList(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

// This small example illustrate how to work with WhiteList function.
func ExampleWhiteList() {
	// Remove all characters from string ignoring characters between "a" and "z"
	println(WhiteList("a3a43a5a4a3a2a23a4a5a4a3a4", "a-z") == "aaaaaaaaaaaa")
}

func Test_BlackList(t *testing.T) {
	tests_1 := []string{"abcdef", "aaaaaaaaaabbbbbbbbbb", "a1b2c3", "   "}
	tests_2 := []string{"abc", "abc", "abc", "abc"}
	expected := []string{"def", "", "123", "   "}
	for i := 0; i < len(tests_1); i++ {
		result := BlackList(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_StripLow(t *testing.T) {
	tests_1 := []string{"foo\x00", "\x7Ffoo\x02", "\x01\x09", "foo\x0A\x0D", "perch\u00e9", "\u20ac",
		"\u2206\x0A", "foo\x0A\x0D", "\x03foo\x0A\x0D"}
	tests_2 := []bool{false, false, false, false, false, false, false, true, true}
	expected := []string{"foo", "foo", "", "foo", "perch\u00e9", "\u20ac", "\u2206", "foo\x0A\x0D", "foo\x0A\x0D"}
	for i := 0; i < len(tests_1); i++ {
		result := StripLow(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_ReplacePattern(t *testing.T) {
	tests_1 := []string{"ab123ba", "abacaba", "httpftp://github.comio", "aaaaaaaaaa", "http123123ftp://git534543hub.comio"}
	tests_2 := []string{"[0-9]+", "[0-9]+", "(ftp|io)", "a", "(ftp|io|[0-9]+)"}
	tests_3 := []string{"aca", "aca", "", "", ""}
	expected := []string{"abacaba", "abacaba", "http://github.com", "", "http://github.com"}
	for i := 0; i < len(tests_1); i++ {
		result := ReplacePattern(tests_1[i], tests_2[i], tests_3[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

// This small example illustrate how to work with ReplacePattern function.
func ExampleReplacePattern() {
	// Replace in "http123123ftp://git534543hub.comio" following (pattern "(ftp|io|[0-9]+)"):
	// - Sequence "ftp".
	// - Sequence "io".
	// - Sequence of digits.
	// with empty string.
	println(ReplacePattern("http123123ftp://git534543hub.comio", "(ftp|io|[0-9]+)", "") == "http://github.com")
}

func Test_Escape(t *testing.T) {
	tests := []string{`<img alt="foo&bar">`}
	expected := []string{"&lt;img alt=&quot;foo&amp;bar&quot;&gt;"}
	for i := 0; i < len(tests); i++ {
		res := Escape(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func Test_UnderscoreToCamelCase(t *testing.T) {
	tests := []string{"a_b_c", "my_func", "1ab_cd"}
	expected := []string{"ABC", "MyFunc", "1abCd"}
	for i := 0; i < len(tests); i++ {
		res := UnderscoreToCamelCase(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func Test_CamelCaseToUnderscore(t *testing.T) {
	tests := []string{"MyFunc", "ABC", "1B"}
	expected := []string{"my_func", "a_b_c", "1_b"}
	for i := 0; i < len(tests); i++ {
		res := CamelCaseToUnderscore(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func Test_Reverse(t *testing.T) {
	tests := []string{"abc", "ｶﾀｶﾅ"}
	expected := []string{"cba", "ﾅｶﾀｶ"}
	for i := 0; i < len(tests); i++ {
		res := Reverse(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func Test_GetLines(t *testing.T) {
	tests := []string{"abc", "a\nb\nc"}
	expected := []([]string){{"abc"}, {"a", "b", "c"}}
	for i := 0; i < len(tests); i++ {
		res := GetLines(tests[i])
		for j := 0; j < len(res); j++ {
			if res[j] != expected[i][j] {
				t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
				t.FailNow()
			}
		}
	}
}

func Test_GetLine(t *testing.T) {
	tests_1 := []string{"abc", "a\nb\nc", "abc", "abacaba\n", "abc"}
	tests_2 := []int{0, 0, -1, 1, 3}
	expected := []string{"abc", "a", "", "", ""}
	for i := 0; i < len(tests_1); i++ {
		res, _ := GetLine(tests_1[i], tests_2[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}
