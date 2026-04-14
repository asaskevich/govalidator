package govalidator

import "testing"

func BenchmarkContains(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Contains("a0b01c012deffghijklmnopqrstu0123456vwxyz", "0123456789")
	}
}

func BenchmarkMatches(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Matches("alfkjl12309fdjldfsa209jlksdfjLAKJjs9uJH234", "[\\w\\d]+")
	}
}

func BenchmarkLeftTrim(b *testing.B) {
	s := "...hello"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = LeftTrim(s, ".")
	}
}

func BenchmarkRightTrim(b *testing.B) {
	s := "hello..."
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = RightTrim(s, ".")
	}
}

func BenchmarkTrim(b *testing.B) {
	s := "   text   "
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Trim(s, "")
	}
}

func BenchmarkWhiteList(b *testing.B) {
	s := "a3a43a5a4a3a2a23a4a5a4a3a4"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = WhiteList(s, "a-z")
	}
}

func BenchmarkBlackList(b *testing.B) {
	s := "a1b2c3d4"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = BlackList(s, "0-9")
	}
}

func BenchmarkStripLow(b *testing.B) {
	s := "Hello\x00\nWorld"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = StripLow(s, true)
	}
}

func BenchmarkReplacePattern(b *testing.B) {
	s := "http123123ftp://git534543hub.comio"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = ReplacePattern(s, "(ftp|io|[0-9]+)", "")
	}
}

func BenchmarkEscape(b *testing.B) {
	s := "<div>hello & \"world\"</div>"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Escape(s)
	}
}

func BenchmarkUnderscoreToCamelCase(b *testing.B) {
	s := "my_func_name"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = UnderscoreToCamelCase(s)
	}
}

func BenchmarkCamelCaseToUnderscore(b *testing.B) {
	s := "MyHTTPServer"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = CamelCaseToUnderscore(s)
	}
}

func BenchmarkReverse(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Reverse(s)
	}
}

func BenchmarkGetLines(b *testing.B) {
	s := "line1\nline2\nline3\nline4\nline5"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = GetLines(s)
	}
}

func BenchmarkGetLine(b *testing.B) {
	s := "line1\nline2\nline3"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = GetLine(s, 1)
	}
}

func BenchmarkRemoveTags(b *testing.B) {
	s := "<p>Hello <b>world</b></p>"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = RemoveTags(s)
	}
}

func BenchmarkSafeFileName(b *testing.B) {
	s := "My Report 2026/04.txt"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = SafeFileName(s)
	}
}

func BenchmarkNormalizeEmail(b *testing.B) {
	s := "Some.One+tag@GoogleMail.com"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = NormalizeEmail(s)
	}
}

func BenchmarkTruncate(b *testing.B) {
	s := "The quick brown fox jumps over the lazy dog"
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Truncate(s, 15, "...")
	}
}

func BenchmarkPadLeft(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = PadLeft("go", "0", 8)
	}
}

func BenchmarkPadRight(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = PadRight("go", "0", 8)
	}
}

func BenchmarkPadBoth(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = PadBoth("go", "0", 8)
	}
}

func BenchmarkTruncatingErrorf(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = TruncatingErrorf("Value: %s", "first", "second")
	}
}
