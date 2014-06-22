package govalidator

import "testing"

func Test_IsAlpha(t *testing.T) {
	if IsAlpha("") {
		t.FailNow()
	}
	if IsAlpha("   foo   ") {
		t.FailNow()
	}
	if IsAlpha("abc1") {
		t.FailNow()
	}
	if !IsAlpha("abc") {
		t.FailNow()
	}
	if !IsAlpha("ABC") {
		t.FailNow()
	}
	if !IsAlpha("FoObAr") {
		t.FailNow()
	}
}

func Test_IsAlphanumeric(t *testing.T) {
	if IsAlphanumeric("foo ") {
		t.FailNow()
	}
	if IsAlphanumeric("abc!!!") {
		t.FailNow()
	}
	if !IsAlphanumeric("abc123") {
		t.FailNow()
	}
	if !IsAlphanumeric("ABC111") {
		t.FailNow()
	}
}

func Test_IsNumeric(t *testing.T) {
	if !IsNumeric("123") {
		t.FailNow()
	}
	if !IsNumeric("0123") {
		t.FailNow()
	}
	if !IsNumeric("-00123") {
		t.FailNow()
	}
	if !IsNumeric("0") {
		t.FailNow()
	}
	if !IsNumeric("-0") {
		t.FailNow()
	}
	if IsNumeric("123.123") {
		t.FailNow()
	}
	if IsNumeric(" ") {
		t.FailNow()
	}
	if IsNumeric(".") {
		t.FailNow()
	}
}

func Test_IsLowerCase(t *testing.T) {
	if !IsLowerCase("abc123") {
		t.FailNow()
	}
	if !IsLowerCase("abc") {
		t.FailNow()
	}
	if !IsLowerCase("tr竪s 端ber") {
		t.FailNow()
	}
	if IsLowerCase("fooBar") {
		t.FailNow()
	}
	if IsLowerCase("123ABC") {
		t.FailNow()
	}
}

func Test_IsUpperCase(t *testing.T) {
	if !IsUpperCase("ABC123") {
		t.FailNow()
	}
	if !IsUpperCase("ABC") {
		t.FailNow()
	}
	if !IsUpperCase("S T R") {
		t.FailNow()
	}
	if IsUpperCase("fooBar") {
		t.FailNow()
	}
	if IsUpperCase("abacaba123") {
		t.FailNow()
	}
}

func Test_IsInt(t *testing.T) {
	if !IsInt("123") {
		t.FailNow()
	}
	if !IsInt("0") {
		t.FailNow()
	}
	if !IsInt("-0") {
		t.FailNow()
	}
	if !IsInt("-0") {
		t.FailNow()
	}
	if IsInt("01") {
		t.FailNow()
	}
	if IsInt("123.123") {
		t.FailNow()
	}
	if IsInt(" ") {
		t.FailNow()
	}
	if IsInt("000") {
		t.FailNow()
	}
}

func Test_IsEmail(t *testing.T) {
	if !IsEmail("foo@bar.com") {
		t.FailNow()
	}
	if !IsEmail("x@x.x") {
		t.FailNow()
	}
	if !IsEmail("foo@bar.com.au") {
		t.FailNow()
	}
	if !IsEmail("foo+bar@bar.com") {
		t.FailNow()
	}
	if IsEmail("invalidemail@") {
		t.FailNow()
	}
	if IsEmail("invalid.com") {
		t.FailNow()
	}
	if IsEmail("@invalid.com") {
		t.FailNow()
	}
}

func Test_IsURL(t *testing.T) {
	if !IsURL("http://foobar.com") {
		t.FailNow()
	}
	if !IsURL("https://foobar.com") {
		t.FailNow()
	}
	if !IsURL("foobar.com") {
		t.FailNow()
	}
	if !IsURL("http://foobar.org/") {
		t.FailNow()
	}
	if !IsURL("http://foobar.org:8080/") {
		t.FailNow()
	}
	if !IsURL("ftp://foobar.ru/") {
		t.FailNow()
	}
	if !IsURL("http://user:pass@www.foobar.com/") {
		t.FailNow()
	}
	if !IsURL("http://127.0.0.1/") {
		t.FailNow()
	}
	if !IsURL("http://duckduckgo.com/?q=%2F") {
		t.FailNow()
	}
	if !IsURL("http://localhost:3000/") {
		t.FailNow()
	}
	if !IsURL("http://foobar.com/?foo=bar#baz=qux") {
		t.FailNow()
	}
	if !IsURL("http://foobar.com?foo=bar") {
		t.FailNow()
	}
	if !IsURL("http://www.xn--froschgrn-x9a.net/") {
		t.FailNow()
	}
	if IsURL("") {
		t.FailNow()
	}
	if IsURL("xyz://foobar.com") {
		t.FailNow()
	}
	if IsURL("invalid.") {
		t.FailNow()
	}
	if IsURL(".com") {
		t.FailNow()
	}
	if IsURL("rtmp://foobar.com") {
		t.FailNow()
	}
	if IsURL("http://www.foo_bar.com/") {
		t.FailNow()
	}
}

func Test_IsFloat(t *testing.T) {
	if IsFloat("") {
		t.FailNow()
	}
	if IsFloat("  ") {
		t.FailNow()
	}
	if IsFloat("-.123") {
		t.FailNow()
	}
	if IsFloat("abacaba") {
		t.FailNow()
	}
	if !IsFloat("123") {
		t.FailNow()
	}
	if !IsFloat("123.") {
		t.FailNow()
	}
	if !IsFloat("123.123") {
		t.FailNow()
	}
	if !IsFloat("-123.123") {
		t.FailNow()
	}
	if !IsFloat("0.123") {
		t.FailNow()
	}
	if !IsFloat("-0.123") {
		t.FailNow()
	}
	if !IsFloat(".0") {
		t.FailNow()
	}
	if !IsFloat("01.123") {
		t.FailNow()
	}
	if !IsFloat("-0.22250738585072011e-307") {
		t.FailNow()
	}
}

func Test_IsHexadecimal(t *testing.T) {
	if IsHexadecimal("abcdefg") {
		t.FailNow()
	}
	if IsHexadecimal("") {
		t.FailNow()
	}
	if IsHexadecimal("..") {
		t.FailNow()
	}
	if !IsHexadecimal("deadBEEF") {
		t.FailNow()
	}
	if !IsHexadecimal("ff0044") {
		t.FailNow()
	}
}

func Test_IsHexcolor(t *testing.T) {
	if IsHexcolor("#ff") {
		t.FailNow()
	}
	if IsHexcolor("fff0") {
		t.FailNow()
	}
	if IsHexcolor("#ff12FG") {
		t.FailNow()
	}
	if !IsHexcolor("#ff0034") {
		t.FailNow()
	}
	if !IsHexcolor("CCccCC") {
		t.FailNow()
	}
	if !IsHexcolor("fff") {
		t.FailNow()
	}
	if !IsHexcolor("#f00") {
		t.FailNow()
	}
}

func Test_IsNull(t *testing.T) {
	if IsNull("abacaba") || !IsNull("") {
		t.FailNow()
	}
}

func Test_IsDivisibleBy(t *testing.T) {
	if !IsDivisibleBy("4", "2") {
		t.FailNow()
	}
	if !IsDivisibleBy("100", "10") {
		t.FailNow()
	}
	if !IsDivisibleBy("", "1") {
		t.FailNow()
	}
	if IsDivisibleBy("123", "foo") {
		t.FailNow()
	}
	if IsDivisibleBy("123", "0") {
		t.FailNow()
	}
}

func Test_IsByteLength(t *testing.T) {
	if IsByteLength("abacaba", 100, -1) {
		t.FailNow()
	}
	if IsByteLength("abacaba", 1, 3) {
		t.FailNow()
	}
	if !IsByteLength("abacaba", 1, 7) {
		t.FailNow()
	}
	if !IsByteLength("abacaba", 0, 8) {
		t.FailNow()
	}
	if IsByteLength("\ufff0", 1, 1) {
		t.FailNow()
	}
}
