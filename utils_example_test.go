package govalidator

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

func ExampleReplacePattern() {
	// Replace in "http123123ftp://git534543hub.comio" following (pattern "(ftp|io|[0-9]+)"):
	// - Sequence "ftp".
	// - Sequence "io".
	// - Sequence of digits.
	// with empty string.
	println(ReplacePattern("http123123ftp://git534543hub.comio", "(ftp|io|[0-9]+)", "") == "http://github.com")
}
