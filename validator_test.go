package govalidator

import (
	"strings"
	"testing"
)

func TestIsAlpha(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", true},
		{"소주", false},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", false}, //UTF-8(ASCII): 0
		{"123", false},
		{"0123", false},
		{"-00123", false},
		{"0", false},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := IsAlpha(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsAlpha(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFLetter(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"", true},
		{"abc", true},
		{"소주", true},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", true},
		{"소", true},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", false}, //UTF-8(ASCII): 0
		{"123", false},
		{"0123", false},
		{"-00123", false},
		{"0", false},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", true},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := IsUTFLetter(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFLetter(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc123", true},
		{"ABC111", true},
		{"abc1", true},
		{"abc〩", false},
		{"abc", true},
		{"소주", false},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := IsAlphanumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsAlphanumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFLetterNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", true},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", true},
		{"abc〩", true},
		{"abc", true},
		{"소주", true},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", true},
		{"소", true},
		{"달기&Co.", false},
		{"〩Hours", true},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", true},
		{"〥〩", true},
		{"모자", true},
		{"ix", true},
		{"۳۵۶۰", true},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := IsUTFLetterNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFLetterNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", true},
		{"+00123", true},
		{"0", true},
		{"-0", true},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", true},
	}
	for _, test := range tests {
		actual := IsNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", true},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", true},
		{"0", true},
		{"-0", true},
		{"--0", false},
		{"-0-", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", true},
		{"-1¾", true},
		{"1¾", true},
		{"〥〩", true},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", true},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", true},
	}
	for _, test := range tests {
		actual := IsUTFNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUTFDigit(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{

		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", true},
		{"0", true},
		{"-0", true},
		{"--0", false},
		{"-0-", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", false},
		{"1483920", true},
		{"", true},
		{"۳۵۶۰", true},
		{"-29", true},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", true},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", true},
	}
	for _, test := range tests {
		actual := IsUTFDigit(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUTFDigit(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsLowerCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc123", true},
		{"abc", true},
		{"a b c", true},
		{"abcß", true},
		{"abcẞ", false},
		{"ABCẞ", false},
		{"tr竪s 端ber", true},
		{"fooBar", false},
		{"123ABC", false},
		{"ABC123", false},
		{"ABC", false},
		{"S T R", false},
		{"fooBar", false},
		{"abacaba123", true},
	}
	for _, test := range tests {
		actual := IsLowerCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsLowerCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUpperCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc123", false},
		{"abc", false},
		{"a b c", false},
		{"abcß", false},
		{"abcẞ", false},
		{"ABCẞ", true},
		{"tr竪s 端ber", false},
		{"fooBar", false},
		{"123ABC", true},
		{"ABC123", true},
		{"ABC", true},
		{"S T R", true},
		{"fooBar", false},
		{"abacaba123", false},
	}
	for _, test := range tests {
		actual := IsUpperCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUpperCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsInt(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"-2147483648", true},          //Signed 32 Bit Min Int
		{"2147483647", true},           //Signed 32 Bit Max Int
		{"-2147483649", true},          //Signed 32 Bit Min Int - 1
		{"2147483648", true},           //Signed 32 Bit Max Int + 1
		{"4294967295", true},           //Unsigned 32 Bit Max Int
		{"4294967296", true},           //Unsigned 32 Bit Max Int + 1
		{"-9223372036854775808", true}, //Signed 64 Bit Min Int
		{"9223372036854775807", true},  //Signed 64 Bit Max Int
		{"-9223372036854775809", true}, //Signed 64 Bit Min Int - 1
		{"9223372036854775808", true},  //Signed 64 Bit Max Int + 1
		{"18446744073709551615", true}, //Unsigned 64 Bit Max Int
		{"18446744073709551616", true}, //Unsigned 64 Bit Max Int + 1
		{"", true},
		{"123", true},
		{"0", true},
		{"-0", true},
		{"+0", true},
		{"01", false},
		{"123.123", false},
		{" ", false},
		{"000", false},
	}
	for _, test := range tests {
		actual := IsInt(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsInt(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		actual := IsEmail(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsEmail(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"", false},
		{"xyz://foobar.com", false},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", false},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", false},
		{"http://www.foo---bar.com/", false},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", false},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
		{"http://foo^bar.org", false},
		{"http://foo&*bar.org", false},
		{"http://foo&bar.org", false},
		{"http://foo bar.org", false},
	}
	for _, test := range tests {
		actual := IsURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsRequestURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"", false},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
	}
	for _, test := range tests {
		actual := IsRequestURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsRequestURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsRequestURI(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"irc://#channel@network", true},
		{"/abs/test/dir", true},
		{"./rel/test/dir", false},
	}
	for _, test := range tests {
		actual := IsRequestURI(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsRequestURI(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsFloat(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"  ", false},
		{"-.123", false},
		{"abacaba", false},
		{"1f", false},
		{"-1f", false},
		{"+1f", false},
		{"123", true},
		{"123.", true},
		{"123.123", true},
		{"-123.123", true},
		{"+123.123", true},
		{"0.123", true},
		{"-0.123", true},
		{"+0.123", true},
		{".0", true},
		{"01.123", true},
		{"-0.22250738585072011e-307", true},
		{"+0.22250738585072011e-307", true},
	}
	for _, test := range tests {
		actual := IsFloat(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsFloat(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsHexadecimal(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abcdefg", false},
		{"", false},
		{"..", false},
		{"deadBEEF", true},
		{"ff0044", true},
	}
	for _, test := range tests {
		actual := IsHexadecimal(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsHexadecimal(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsHexcolor(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"#ff", false},
		{"fff0", false},
		{"#ff12FG", false},
		{"CCccCC", true},
		{"fff", true},
		{"#f00", true},
	}
	for _, test := range tests {
		actual := IsHexcolor(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsHexcolor(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsRGBcolor(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"rgb(0,31,255)", true},
		{"rgb(1,349,275)", false},
		{"rgb(01,31,255)", false},
		{"rgb(0.6,31,255)", false},
		{"rgba(0,31,255)", false},
		{"rgb(0,  31, 255)", true},
	}
	for _, test := range tests {
		actual := IsRGBcolor(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsRGBcolor(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsNull(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abacaba", false},
		{"", true},
	}
	for _, test := range tests {
		actual := IsNull(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNull(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsDivisibleBy(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param1   string
		param2   string
		expected bool
	}{
		{"4", "2", true},
		{"100", "10", true},
		{"", "1", true},
		{"123", "foo", false},
		{"123", "0", false},
	}
	for _, test := range tests {
		actual := IsDivisibleBy(test.param1, test.param2)
		if actual != test.expected {
			t.Errorf("Expected IsDivisibleBy(%q, %q) to be %v, got %v", test.param1, test.param2, test.expected, actual)
		}
	}
}

// This small example illustrate how to work with IsDivisibleBy function.
func ExampleIsDivisibleBy() {
	println("1024 is divisible by 64: ", IsDivisibleBy("1024", "64"))
}

func TestIsByteLength(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param1   string
		param2   int
		param3   int
		expected bool
	}{
		{"abacaba", 100, -1, false},
		{"abacaba", 1, 3, false},
		{"abacaba", 1, 7, true},
		{"abacaba", 0, 8, true},
		{"\ufff0", 1, 1, false},
	}
	for _, test := range tests {
		actual := IsByteLength(test.param1, test.param2, test.param3)
		if actual != test.expected {
			t.Errorf("Expected IsByteLength(%q, %q, %q) to be %v, got %v", test.param1, test.param2, test.param3, test.expected, actual)
		}
	}
}

func TestIsJSON(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"145", true},
		{"asdf", false},
		{"123:f00", false},
		{"{\"Name\":\"Alice\",\"Body\":\"Hello\",\"Time\":1294706395881547000}", true},
		{"{}", true},
		{"{\"Key\":{\"Key\":{\"Key\":123}}}", true},
		{"[]", true},
		{"null", true},
	}
	for _, test := range tests {
		actual := IsJSON(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsJSON(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsMultibyte(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abc", false},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"あいうえお foobar", true},
		{"test＠example.com", true},
		{"test＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"ｶﾀｶﾅ", true},
	}
	for _, test := range tests {
		actual := IsMultibyte(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsMultibyte(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsASCII(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"", true},
	}
	for _, test := range tests {
		actual := IsASCII(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsASCII(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsPrintableASCII(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"newline\n", false},
		{"\x19test\x7F", false},
	}
	for _, test := range tests {
		actual := IsPrintableASCII(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsPrintableASCII(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsFullWidth(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc", false},
		{"abc123", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"３ー０　ａ＠ｃｏｍ", true},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"Good＝Parts", true},
		{"", true},
	}
	for _, test := range tests {
		actual := IsFullWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsFullWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsHalfWidth(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"あいうえお", false},
		{"００１１", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", true},
		{"l-btn_02--active", true},
		{"abc123い", true},
		{"ｶﾀｶﾅﾞﾬ￩", true},
		{"", true},
	}
	for _, test := range tests {
		actual := IsHalfWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsHalfWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsVariableWidth(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ひらがなカタカナ漢字ABCDE", true},
		{"３ー０123", true},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"", true},
		{"Good＝Parts", true},
		{"abc", false},
		{"abc123", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", false},
		{"ひらがな・カタカナ、．漢字", false},
		{"１２３４５６", false},
		{"ｶﾀｶﾅﾞﾬ", false},
	}
	for _, test := range tests {
		actual := IsVariableWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsVariableWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsUUID(t *testing.T) {
	t.Parallel()

	// Tests without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3xxx", false},
		{"a987fbc94bed3078cf079141ba07c9f3", false},
		{"934859", false},
		{"987fbc9-4bed-3078-cf07a-9141ba07c9f3", false},
		{"aaaaaaaa-1111-1111-aaag-111111111111", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}
	for _, test := range tests {
		actual := IsUUID(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUUID(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// UUID ver. 3
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"412452646", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-4078-8f07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}
	for _, test := range tests {
		actual := IsUUIDv3(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUUIDv3(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// UUID ver. 4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-5078-af07-9141ba07c9f3", false},
		{"934859", false},
		{"57b73598-8764-4ad0-a76a-679bb6640eb1", true},
		{"625e63f3-58f5-40b7-83a1-a72ad31acffb", true},
	}
	for _, test := range tests {
		actual := IsUUIDv4(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUUIDv4(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// UUID ver. 5
	tests = []struct {
		param    string
		expected bool
	}{

		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"9c858901-8a57-4791-81fe-4c455b099bc9", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"987fbc97-4bed-5078-af07-9141ba07c9f3", true},
		{"987fbc97-4bed-5078-9f07-9141ba07c9f3", true},
	}
	for _, test := range tests {
		actual := IsUUIDv5(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUUIDv5(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsCreditCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5398228707871528", false},
		{"375556917985515", true},
		{"36050234196908", true},
		{"4716461583322103", true},
		{"4716-2210-5188-5662", true},
		{"4929 7226 5379 7141", true},
		{"5398228707871527", true},
	}
	for _, test := range tests {
		actual := IsCreditCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsCreditCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsISBN(t *testing.T) {
	t.Parallel()

	// Without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}
	for _, test := range tests {
		actual := IsISBN(test.param, -1)
		if actual != test.expected {
			t.Errorf("Expected IsISBN(%q, -1) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// ISBN 10
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3423214121", false},
		{"978-3836221191", false},
		{"3-423-21412-1", false},
		{"3 423 21412 1", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
	}
	for _, test := range tests {
		actual := IsISBN10(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsISBN10(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// ISBN 13
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3-8362-2119-5", false},
		{"01234567890ab", false},
		{"978 3 8362 2119 0", false},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}
	for _, test := range tests {
		actual := IsISBN13(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsISBN13(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsDataURI(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"data:text/plain;base64,Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
		{"data:image/gif;base64,MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"data:image/png;base64,12345", false},
		{"", false},
		{"data:text,:;base85,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
	}
	for _, test := range tests {
		actual := IsDataURI(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsDataURI(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsBase64(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", true},
		{"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"12345", false},
		{"", false},
		{"Vml2YW11cyBmZXJtZtesting123", false},
	}
	for _, test := range tests {
		actual := IsBase64(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsBase64(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsISO3166Alpha2(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"ABCD", false},
		{"A", false},
		{"AC", false},
		{"AP", false},
		{"GER", false},
		{"NU", true},
		{"DE", true},
		{"JP", true},
		{"JPN", false},
		{"ZWE", false},
		{"GER", false},
		{"DEU", false},
	}
	for _, test := range tests {
		actual := IsISO3166Alpha2(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsISO3166Alpha2(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsISO3166Alpha3(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"ABCD", false},
		{"A", false},
		{"AC", false},
		{"AP", false},
		{"NU", false},
		{"DE", false},
		{"JP", false},
		{"ZWE", true},
		{"JPN", true},
		{"GER", false},
		{"DEU", true},
	}
	for _, test := range tests {
		actual := IsISO3166Alpha3(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsISO3166Alpha3(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsIP(t *testing.T) {
	t.Parallel()

	// Without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsIP(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIP(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", false},
		{"2001:db8:0000:1:1:1:1:1", false},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsIPv4(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIPv4(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv6
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"1.2.3.4", false},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IsIPv6(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIPv6(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsMAC(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"3D:F2:C9:A6:B3:4F", true},
		{"3D-F2-C9-A6-B3:4F", false},
		{"123", false},
		{"", false},
		{"abacaba", false},
	}
	for _, test := range tests {
		actual := IsMAC(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsMAC(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestFilePath(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
		osType   int
	}{
		{"c:\\" + strings.Repeat("a", 32767), true, Win}, //See http://msdn.microsoft.com/en-us/library/aa365247(VS.85).aspx#maxpath
		{"c:\\" + strings.Repeat("a", 32768), false, Win},
		{"c:\\path\\file (x86)\bar", true, Win},
		{"c:\\path\\file", true, Win},
		{"c:\\path\\file:exe", false, Unknown},
		{"C:\\", true, Win},
		{"c:\\path\\file\\", true, Win},
		{"c:/path/file/", false, Unknown},
		{"/path/file/", true, Unix},
		{"/path/file:SAMPLE/", true, Unix},
		{"/path/file:/.txt", true, Unix},
		{"/path", true, Unix},
	}
	for _, test := range tests {
		actual, osType := IsFilePath(test.param)
		if actual != test.expected || osType != test.osType {
			t.Errorf("Expected IsFilePath(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsLatitude(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-90.000", true},
		{"+90", true},
		{"47.1231231", true},
		{"+99.9", false},
		{"108", false},
	}
	for _, test := range tests {
		actual := IsLatitude(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsLatitude(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsLongitude(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-180.000", true},
		{"180.1", false},
		{"+73.234", true},
		{"+382.3811", false},
		{"23.11111111", true},
	}
	for _, test := range tests {
		actual := IsLongitude(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsLongitude(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsSSN(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"00-90-8787", false},
		{"66690-76", false},
		{"191 60 2869", true},
		{"191-60-2869", true},
	}
	for _, test := range tests {
		actual := IsSSN(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsSSN(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIsMongoID(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"507f1f77bcf86cd799439011", true},
		{"507f1f77bcf86cd7994390", false},
		{"507f1f77bcf86cd79943901z", false},
		{"507f1f77bcf86cd799439011 ", false},
		{"", false},
	}
	for _, test := range tests {
		actual := IsMongoID(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsMongoID(%q) to be %v, got %v", test.param, test.expected, actual)
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

type UserValid struct {
	Name     string `valid:"required"`
	Email    string `valid:"required,email"`
	Password string `valid:"required"`
	Age      int    `valid:"required"`
	Home     *Address
	Work     []Address
}

type PrivateStruct struct {
	privateField string `valid:"required,alpha,d_k"`
	NonZero      int
	ListInt      []int
	ListString   []string `valid:"alpha"`
	Work         [2]Address
	Home         Address
	Map          map[string]Address
}

type NegationStruct struct {
	NotInt string `valid:"!int"`
	Int    string `valid:"int"`
}

func TestValidateNegationStruct(t *testing.T) {
	var tests = []struct {
		param    NegationStruct
		expected bool
	}{
		{NegationStruct{"a1", "11"}, true},
		{NegationStruct{"email@email.email", "11"}, true},
		{NegationStruct{"123456----", "11"}, true},
		{NegationStruct{"::1", "11"}, true},
		{NegationStruct{"123.123", "11"}, true},
		{NegationStruct{"a1", "a1"}, false},
		{NegationStruct{"11", "a1"}, false},
		{NegationStruct{"11", "11"}, false},
	}
	for _, test := range tests {
		actual, err := ValidateStruct(test.param)
		if actual != test.expected {
			t.Errorf("Expected ValidateStruct(%q) to be %v, got %v", test.param, test.expected, actual)
			if err != nil {
				t.Errorf("Got Error on ValidateStruct(%q): %s", test.param, err)
			}
		}
	}
}

func TestValidateStruct(t *testing.T) {

	var tests = []struct {
		param    interface{}
		expected bool
	}{
		{User{"John", "john@yahoo.com", "123G#678", 20, &Address{"Street", "123456"}, []Address{Address{"Street", "123456"}, Address{"Street", "123456"}}}, false},
		{User{"John", "john!yahoo.com", "12345678", 20, &Address{"Street", "ABC456D89"}, []Address{Address{"Street", "ABC456D89"}, Address{"Street", "123456"}}},false},
		{User{"John", "", "12345", 0, &Address{"Street", "123456789"}, []Address{Address{"Street", "ABC456D89"}, Address{"Street", "123456"}}},false},
		{UserValid{"John", "john@yahoo.com", "123G#678", 20, &Address{"Street", "123456"}, []Address{Address{"Street", "123456"}, Address{"Street", "123456"}}}, true},
		{UserValid{"John", "john!yahoo.com", "12345678", 20, &Address{"Street", "ABC456D89"}, []Address{Address{"Street", "ABC456D89"}, Address{"Street", "123456"}}},false},
		{UserValid{"John", "", "12345", 0, &Address{"Street", "123456789"}, []Address{Address{"Street", "ABC456D89"}, Address{"Street", "123456"}}},false},
		{nil,true},
		{User{"John", "john@yahoo.com", "123G#678", 0, &Address{"Street", "123456"}, []Address{}},false},
		{"im not a struct", false},
	}
	for _, test := range tests {
		actual, err := ValidateStruct(test.param)
		if actual != test.expected {
			t.Errorf("Expected ValidateStruct(%q) to be %v, got %v", test.param, test.expected, actual)
			if err != nil {
				t.Errorf("Got Error on ValidateStruct(%q): %s", test.param, err)
			}
		}
	}



	TagMap["d_k"] = Validator(func(str string) bool {
		return str == "d_k"
	})
	result, err := ValidateStruct(PrivateStruct{"d_k", 0, []int{1, 2}, []string{"hi", "super"}, [2]Address{Address{"Street", "123456"},
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
