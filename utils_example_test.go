package govalidator

func ExampleContains() {
	println(Contains("Hello, world", "world"))
}

func ExampleMatches() {
	println(Matches("govalidator", "^[a-z]+$"))
}

func ExampleLeftTrim() {
	println(LeftTrim("...hello", ".") == "hello")
}

func ExampleRightTrim() {
	println(RightTrim("hello...", ".") == "hello")
}

func ExampleTrim() {
	// Remove from left and right spaces and "\r", "\n", "\t" characters
	println(Trim("   \r\r\ntext\r   \t\n", "") == "text")
	// Remove from left and right characters that are between "1" and "8".
	// "1-8" is like full list "12345678".
	println(Trim("1234567890987654321", "1-8") == "909")
}

func ExampleWhiteList() {
	// Remove all characters from string ignoring characters between "a" and "z"
	println(WhiteList("a3a43a5a4a3a2a23a4a5a4a3a4", "a-z") == "aaaaaaaaaaaa")
}

func ExampleBlackList() {
	println(BlackList("a1b2c3d4", "0-9") == "abcd")
}

func ExampleStripLow() {
	println(StripLow("Hello\x00\nWorld", true) == "Hello\nWorld")
}

func ExampleReplacePattern() {
	// Replace in "http123123ftp://git534543hub.comio" following (pattern "(ftp|io|[0-9]+)"):
	// - Sequence "ftp".
	// - Sequence "io".
	// - Sequence of digits.
	// with empty string.
	println(ReplacePattern("http123123ftp://git534543hub.comio", "(ftp|io|[0-9]+)", "") == "http://github.com")
}

func ExampleEscape() {
	println(Escape("<div>hello & \"world\"</div>") == "&lt;div&gt;hello &amp; &#34;world&#34;&lt;/div&gt;")
}

func ExampleUnderscoreToCamelCase() {
	println(UnderscoreToCamelCase("my_func_name") == "MyFuncName")
}

func ExampleCamelCaseToUnderscore() {
	println(CamelCaseToUnderscore("MyHTTPServer") == "my_http_server")
}

func ExampleReverse() {
	println(Reverse("abc") == "cba")
}

func ExampleGetLines() {
	_ = GetLines("line1\nline2\nline3") // result = []string{"line1", "line2", "line3"}
}

func ExampleGetLine() {
	_, _ = GetLine("line1\nline2\nline3", 1) // line2, nil
}

func ExampleRemoveTags() {
	println(RemoveTags("<p>Hello <b>world</b></p>") == "Hello world")
}

func ExampleSafeFileName() {
	println(SafeFileName("My Report 2026/04.txt") == "04.txt")
}

func ExampleNormalizeEmail() {
	_, _ = NormalizeEmail("Some.One+tag@GoogleMail.com") // someone@gmail.com, nil
}

func ExampleTruncate() {
	println(Truncate("The quick brown fox jumps", 15, "...") == "The quick brown...")
}

func ExamplePadLeft() {
	println(PadLeft("go", "0", 5) == "000go")
}

func ExamplePadRight() {
	println(PadRight("go", "0", 5) == "go000")
}

func ExamplePadBoth() {
	println(PadBoth("go", "0", 6) == "00go00")
}

func ExampleTruncatingErrorf() {
	_ = TruncatingErrorf("Value: %s", "first", "second") // Value: first
}
