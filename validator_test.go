package govalidator

import "testing"

func TestIsAlpha(t *testing.T) {
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

func TestIsAlphanumeric(t *testing.T) {
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

func TestIsNumeric(t *testing.T) {
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

func TestIsLowerCase(t *testing.T) {
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

func TestIsUpperCase(t *testing.T) {
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

func TestIsInt(t *testing.T) {
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

func TestIsEmail(t *testing.T) {
	tests := []string{"foo@bar.com", "x@x.x", "foo@bar.com.au", "foo+bar@bar.com", "invalidemail@", "invalid.com", "@invalid.com",
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

func TestIsURL(t *testing.T) {
	tests := []string{"http://foobar.com", "https://foobar.com", "foobar.com", "http://foobar.org/", "http://foobar.org:8080/",
		"ftp://foobar.ru/", "http://user:pass@www.foobar.com/", "http://127.0.0.1/", "http://duckduckgo.com/?q=%2F", "http://localhost:3000/",
		"http://foobar.com/?foo=bar#baz=qux", "http://foobar.com?foo=bar", "http://www.xn--froschgrn-x9a.net/",
		"", "xyz://foobar.com", "invalid.", ".com", "rtmp://foobar.com", "http://www.foo_bar.com/", "http://localhost:3000/",
		"http://foobar.com#baz=qux", "http://foobar.com/t$-_.+!*\\'(),", "http://www.foobar.com/~foobar", "http://www.-foobar.com/",
		"http://www.foo---bar.com/"}
	expected := []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, false, false, false,
		false, false, true, true, true, true, true, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsURL(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result, tests[i])
			t.FailNow()
		}
	}
}

func TestIsFloat(t *testing.T) {
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

func TestIsHexadecimal(t *testing.T) {
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

func TestIsHexcolor(t *testing.T) {
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

func TestIsRGBcolor(t *testing.T) {
	tests := []string{"rgb(0,31,255)", "rgb(1,349,275)", "rgb(01,31,255)", "rgb(0.6,31,255)", "rgba(0,31,255)", "rgb(0,  31, 255)"}
	expected := []bool{true, false, false, false, false, true}
	for i := 0; i < len(tests); i++ {
		result := IsRGBcolor(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsNull(t *testing.T) {
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

func TestIsDivisibleBy(t *testing.T) {
	tests1 := []string{"4", "100", "", "123", "123"}
	tests2 := []string{"2", "10", "1", "foo", "0"}
	expected := []bool{true, true, true, false, false}
	for i := 0; i < len(tests1); i++ {
		result := IsDivisibleBy(tests1[i], tests2[i])
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

func TestIsByteLength(t *testing.T) {
	tests1 := []string{"abacaba", "abacaba", "abacaba", "abacaba", "\ufff0"}
	tests2 := []int{100, 1, 1, 0, 1}
	tests3 := []int{-1, 3, 7, 8, 1}
	expected := []bool{false, false, true, true, false}
	for i := 0; i < len(tests1); i++ {
		result := IsByteLength(tests1[i], tests2[i], tests3[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsJSON(t *testing.T) {
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

func TestIsMultibyte(t *testing.T) {
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

func TestIsASCII(t *testing.T) {
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

func TestIsFullWidth(t *testing.T) {
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

func TestIsHalfWidth(t *testing.T) {
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

func TestIsVariableWidth(t *testing.T) {
	tests := []string{"ひらがなカタカナ漢字ABCDE", "３ー０123", "Ｆｶﾀｶﾅﾞﾬ", "Good＝Parts", "abc", "abc123",
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

func TestIsUUID(t *testing.T) {
	// Tests without version
	tests := []string{"", "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "A987FBC9-4BED-3078-CF07-9141BA07C9F3xxx",
		"A987FBC94BED3078CF079141BA07C9F3", "934859", "987FBC9-4BED-3078-CF07A-9141BA07C9F3",
		"AAAAAAAA-1111-1111-AAAG-111111111111", "A987FBC9-4BED-3078-CF07-9141BA07C9F3"}
	expected := []bool{false, false, false, false, false, false, false, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUID(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// UUID ver. 3
	tests = []string{"", "412452646", "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "A987FBC9-4BED-4078-8F07-9141BA07C9F3",
		"A987FBC9-4BED-3078-CF07-9141BA07C9F3"}
	expected = []bool{false, false, false, false, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUIDv3(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// UUID ver. 4
	tests = []string{"", "xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "A987FBC9-4BED-5078-AF07-9141BA07C9F3",
		"934859", "57B73598-8764-4AD0-A76A-679BB6640EB1", "625E63F3-58F5-40B7-83A1-A72AD31ACFFB"}
	expected = []bool{false, false, false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUIDv4(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// UUID ver. 5
	tests = []string{"xxxA987FBC9-4BED-3078-CF07-9141BA07C9F3", "9c858901-8a57-4791-81fe-4c455b099bc9", "A987FBC9-4BED-3078-CF07-9141BA07C9F3",
		"", "987FBC97-4BED-5078-AF07-9141BA07C9F3", "987FBC97-4BED-5078-9F07-9141BA07C9F3"}
	expected = []bool{false, false, false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsUUIDv5(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}

}

func TestIsCreditCard(t *testing.T) {
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

func TestIsISBN(t *testing.T) {
	// ISBN 10
	tests := []string{"", "foo", "3423214121", "978-3836221191", "3-423-21412-1", "3 423 21412 1", "3836221195", "1-61729-085-8",
		"3 423 21412 0", "3 401 01319 X"}
	expected := []bool{false, false, false, false, false, false, true, true, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsISBN10(tests[i])
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
		result := IsISBN13(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// Without version
	tests = []string{"3836221195", "1-61729-085-8", "3 423 21412 0", "3 401 01319 X", "9784873113685", "978-4-87311-368-5",
		"978 3401013190", "978-3-8362-2119-1", "", "foo"}
	expected = []bool{true, true, true, true, true, true, true, true, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsISBN(tests[i], -1)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsDataURI(t *testing.T) {
	tests := []string{"data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=",
		"data:text/plain;base64,Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", "image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==",
		"data:image/gif;base64,MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", "data:image/png;base64,12345", "",
		"data:text,:;base85,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw=="}
	expected := []bool{true, true, false, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsDataURI(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsBase64(t *testing.T) {
	tests := []string{"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=",
		"Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", "U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==",
		"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", "12345", "",
		"Vml2YW11cyBmZXJtZtesting123"}
	expected := []bool{true, true, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsBase64(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsIP(t *testing.T) {
	// IPv4
	tests := []string{"127.0.0.1", "0.0.0.0", "255.255.255.255", "1.2.3.4", "::1", "2001:db8:0000:1:1:1:1:1"}
	expected := []bool{true, true, true, true, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsIPv4(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// IPv6
	tests = []string{"127.0.0.1", "0.0.0.0", "255.255.255.255", "1.2.3.4", "::1", "2001:db8:0000:1:1:1:1:1"}
	expected = []bool{false, false, false, false, true, true}
	for i := 0; i < len(tests); i++ {
		result := IsIPv6(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
	// Without version
	tests = []string{"127.0.0.1", "0.0.0.0", "255.255.255.255", "1.2.3.4", "::1", "2001:db8:0000:1:1:1:1:1", "300.0.0.0"}
	expected = []bool{true, true, true, true, true, true, false}
	for i := 0; i < len(tests); i++ {
		result := IsIP(tests[i], -1)
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsMAC(t *testing.T) {
	tests := []string{"3D:F2:C9:A6:B3:4F", "3D-F2-C9-A6-B3:4F", "123", "", "abacaba"}
	expected := []bool{true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsMAC(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsLatitude(t *testing.T) {
	tests := []string{"-90.000", "+90", "47.1231231", "+99.9", "108"}
	expected := []bool{true, true, true, false, false}
	for i := 0; i < len(tests); i++ {
		result := IsLatitude(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestIsLongitude(t *testing.T) {
	tests := []string{"-180.000", "180.1", "+73.234", "+382.3811", "23.11111111"}
	expected := []bool{true, false, true, false, true}
	for i := 0; i < len(tests); i++ {
		result := IsLongitude(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

type Address struct {
	Street string `valid:"-"`
	Zip    string `json:"zip" valid:"numeric,required"`
}

type User struct {
	Name     string `valid:"required"`
	Email    string `valid:"required,email"`
	Password string `valid:"required"`
	Age      int    `valid:"required,numeric,@#\u0000"`
	Home     *Address
	Work     []Address
}

type PrivateStruct struct {
	privateField string `valid:"required,alpha,d_k"`
	NonZero      int
	Work         [2]Address
	Home         Address
	Map          map[string]Address
}

func TestValidateStruct(t *testing.T) {
	// Valid structure
	user := &User{"John", "john@yahoo.com", "123G#678", 20, &Address{"Street", "123456"}, []Address{Address{"Street", "123456"}, Address{"Street", "123456"}}}
	result, err := ValidateStruct(user)
	if result != true {
		t.Log("Case ", 0, ": expected ", true, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}
	// Not valid
	user = &User{"John", "john!yahoo.com", "12345678", 20, &Address{"Street", "ABC456D89"}, []Address{Address{"Street", "ABC456D89"}, Address{"Street", "123456"}}}
	result, err = ValidateStruct(user)
	if result == true {
		t.Log("Case ", 1, ": expected ", false, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}
	user = &User{"John", "", "12345", 0, &Address{"Street", "123456789"}, []Address{Address{"Street", "ABC456D89"}, Address{"Street", "123456"}}}
	result, err = ValidateStruct(user)
	if result == true {
		t.Log("Case ", 2, ": expected ", false, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}
	result, err = ValidateStruct(nil)
	if result != true {
		t.Log("Case ", 3, ": expected ", true, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}
	user = &User{"John", "john@yahoo.com", "123G#678", 0, &Address{"Street", "123456"}, []Address{}}
	result, err = ValidateStruct(user)
	if result != true {
		t.Log("Case ", 4, ": expected ", true, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}
	result, err = ValidateStruct("im not a struct")
	if result == true {
		t.Log("Case ", 5, ": expected ", false, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}

	TagMap["d_k"] = Validator(func(str string) bool {
		return str == "d_k"
	})
	result, err = ValidateStruct(PrivateStruct{"d_k", 0, [2]Address{Address{"Street", "123456"},
		Address{"Street", "123456"}}, Address{"Street", "123456"}, map[string]Address{"address": Address{"Street", "123456"}}})
	if result != true {
		t.Log("Case ", 6, ": expected ", true, " when result is ", result)
		t.Error(err)
		t.FailNow()
	}
}

func ExampleValidateStruct() {
	type Post struct {
		Title    string `valid:"alphanum,required"`
		Message  string `valid:"duck,ascii"`
		AuthorIP string `valid:"ipv4"`
	}
	post := &Post{"My Example Post", "duck", "123.234.54.3"}

	//Add your own struct validation tags
	TagMap["duck"] = Validator(func(str string) bool {
		return str == "duck"
	})

	result, err := ValidateStruct(post)
	if err != nil {
		println("error: " + err.Error())
	}
	println(result)
}
