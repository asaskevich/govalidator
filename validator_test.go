package govalidator

import "testing"

func Test_IsAlpha(t *testing.T) {
	tests := []string{"", "   fooo   ", "abc1", "abc", "ABC", "FoObAr"}
	expected := []bool{false, false, false, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsAlpha(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsAlphanumeric(t *testing.T) {
	tests := []string{"foo ", "abc!!!", "abc123", "ABC111"}
	expected := []bool{false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsAlphanumeric(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsNumeric(t *testing.T) {
	tests := []string{"123", "0123", "-00123", "0", "-0", "123.123", " ", "."}
	expected := []bool{true, true, true, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsNumeric(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsLowerCase(t *testing.T) {
	tests := []string{"abc123", "abc", "tr竪s 端ber", "fooBar", "123ABC"}
	expected := []bool{true, true, true, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsLowerCase(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsUpperCase(t *testing.T) {
	tests := []string{"ABC123", "ABC", "S T R", "fooBar", "abacaba123"}
	expected := []bool{true, true, true, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsUpperCase(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsInt(t *testing.T) {
	tests := []string{"123", "0", "-0", "01", "123.123", " ", "000"}
	expected := []bool{true, true, true, false, false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsInt(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsEmail(t *testing.T) {
	tests := []string{"foo@bar.com", "x@x.x" , "foo@bar.com.au", "foo+bar@bar.com", "invalidemail@", "invalid.com", "@invalid.com",
		"test|123@m端ller.com", "hans@m端ller.com", "hans.m端ller@test.com"}
	expected := []bool{true, true, true, true, false, false, false, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsEmail(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsURL(t *testing.T) {
	tests := []string{"http://foobar.com", "https://foobar.com", "foobar.com", "http://foobar.org/", "http://foobar.org:8080/",
		"ftp://foobar.ru/", "http://user:pass@www.foobar.com/", "http://127.0.0.1/", "http://duckduckgo.com/?q=%2F", "http://localhost:3000/",
		"http://foobar.com/?foo=bar#baz=qux", "http://foobar.com?foo=bar", "http://www.xn--froschgrn-x9a.net/",
		"", "xyz://foobar.com", "invalid.", ".com", "rtmp://foobar.com", "http://www.foo_bar.com/"}
	expected := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false,
		false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsURL(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsFloat(t *testing.T) {
	tests := []string{"", "  ", "-.123", "abacaba", "123", "123.", "123.123", "-123.123", "0.123", "-0.123", ".0",
		"01.123", "-0.22250738585072011e-307"}
	expected := []bool{false, false, false, false, true, true, true, true, true, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsFloat(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsHexadecimal(t *testing.T) {
	tests := []string{"abcdefg", "", "..", "deadBEEF", "ff0044"}
	expected := []bool{false, false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsHexadecimal(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsHexcolor(t *testing.T) {
	tests := []string{"#ff", "fff0", "#ff12FG", "CCccCC", "fff", "#f00"}
	expected := []bool{false, false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsHexcolor(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsNull(t *testing.T) {
	tests := []string{"abacaba", ""}
	expected := []bool{false, true}
	for i := 0; i < len(tests); i++ {
		result := IsNull(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsDivisibleBy(t *testing.T) {
	tests_1 := []string{"4", "100", "", "123", "123"}
	tests_2 := []string{"2", "10", "1", "foo", "0"}
	expected := []bool{true, true, true, false, false}
	for i := 0; i < len(tests_1); i++ {
		result := IsDivisibleBy(tests_1[i], tests_2[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

// This small example illustrate how to work with IsDivisibleBy function.
func ExampleIsDivisibleBy() {
	println("1024 is divisible by 64: ", IsDivisibleBy("1024", "64"))
}

func Test_IsByteLength(t *testing.T) {
	tests_1 := []string{"abacaba", "abacaba", "abacaba", "abacaba", "\ufff0"}
	tests_2 := []int{100, 1, 1, 0, 1}
	tests_3 := []int{-1, 3, 7, 8, 1}
	expected := []bool{false, false, true, true, false}
	for i := 0; i < len(tests_1); i++ {
		result := IsByteLength(tests_1[i], tests_2[i], tests_3[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsJSON(t *testing.T) {
	tests := []string{"", "145", "asdf", "123:f00", "{\"Name\":\"Alice\",\"Body\":\"Hello\",\"Time\":1294706395881547000}",
		"{}", "{\"Key\":{\"Key\":{\"Key\":123}}}"}
	expected := []bool{false, false, false, false, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsJSON(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsMultibyte(t *testing.T) {
	tests := []string{"abc", "123", "<>@;.-=", "ひらがな・カタカナ、．漢字", "あいうえお foobar", "test＠example.com",
		"test＠example.com", "1234abcDEｘｙｚ", "ｶﾀｶﾅ"}
	expected := []bool{false, false, false, true, true, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsMultibyte(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsASCII(t *testing.T) {
	tests := []string{"ｆｏｏbar", "ｘｙｚ０９８", "１２３456", "ｶﾀｶﾅ", "foobar", "0987654321", "test@example.com", "1234abcDEF"}
	expected := []bool{false, false, false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsASCII(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsFullWidth(t *testing.T) {
	tests := []string{"abc", "abc123", "!\"#$%&()<>/+=-_? ~^|.,@`{}[]", "ひらがな・カタカナ、．漢字", "３ー０　ａ＠ｃｏｍ", "Ｆｶﾀｶﾅﾞﾬ", "Good＝Parts"}
	expected := []bool{false, false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsFullWidth(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsHalfWidth(t *testing.T) {
	tests := []string{"あいうえお", "００１１", "!\"#$%&()<>/+=-_? ~^|.,@`{}[]", "l-btn_02--active", "abc123い", "ｶﾀｶﾅﾞﾬ￩"}
	expected := []bool{false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsHalfWidth(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsVariableWidth(t *testing.T) {
	tests := []string{"ひらがなカタカナ漢字ABCDE", "３ー０123", "Ｆｶﾀｶﾅﾞﾬ", "Good＝Parts" , "abc", "abc123",
		"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", "ひらがな・カタカナ、．漢字", "１２３４５６", "ｶﾀｶﾅﾞﾬ"}
	expected := []bool{true, true, true, true, false, false, false, false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsVariableWidth(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsUUID(t *testing.T) {
	// Tests without version
	tests := []string{"", "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "A987FBC9-4BED-3078-CF07-9141BA07C9F3xxx",
		"A987FBC94BED3078CF079141BA07C9F3", "934859", "987FBC9-4BED-3078-CF07A-9141BA07C9F3",
		"AAAAAAAA-1111-1111-AAAG-111111111111", "A987FBC9-4BED-3078-CF07-9141BA07C9F3", }
	expected := []bool{false, false, false, false, false, false, false, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUID(tests[i], 0)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// UUID ver. 3
	tests = []string{"", "412452646", "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "A987FBC9-4BED-4078-8F07-9141BA07C9F3",
		"A987FBC9-4BED-3078-CF07-9141BA07C9F3", }
	expected = []bool{false, false, false, false, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUID(tests[i], 3)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// UUID ver. 4
	tests = []string{"", "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "A987FBC9-4BED-5078-AF07-9141BA07C9F3",
		"934859", "57B73598-8764-4AD0-A76A-679BB6640EB1", "625E63F3-58F5-40B7-83A1-A72AD31ACFFB", }
	expected = []bool{false, false, false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUID(tests[i], 4)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// UUID ver. 5
	tests = []string{"xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "9c858901-8a57-4791-81fe-4c455b099bc9", "A987FBC9-4BED-3078-CF07-9141BA07C9F3",
		"", "987FBC97-4BED-5078-AF07-9141BA07C9F3", "987FBC97-4BED-5078-9F07-9141BA07C9F3", }
	expected = []bool{false, false, false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUID(tests[i], 5)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// Wrong version
	tests = []string{""}
	expected = []bool{false}
	for i := 0; i < len(tests); i++ {
		result := IsUUID(tests[i], -1)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsCreditCard(t *testing.T) {
	tests := []string{"foo", "5398228707871528", "375556917985515", "36050234196908", "4716461583322103", "4716-2210-5188-5662",
		"4929 7226 5379 7141", "5398228707871527"}
	expected := []bool{false, false, true, true, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsCreditCard(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func Test_IsISBN(t *testing.T) {
	// ISBN 10
	tests := []string{"", "foo", "3423214121", "978-3836221191", "3-423-21412-1", "3 423 21412 1", "3836221195", "1-61729-085-8",
		"3 423 21412 0", "3 401 01319 X"}
	expected := []bool{false, false, false, false, false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsISBN(tests[i], 10)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// ISBN 13
	tests = []string{"", "3-8362-2119-5", "01234567890ab", "978 3 8362 2119 0", "9784873113685", "978-4-87311-368-5",
		"978 3401013190", "978-3-8362-2119-1"}
	expected = []bool{false, false, false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsISBN(tests[i], 13)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}
