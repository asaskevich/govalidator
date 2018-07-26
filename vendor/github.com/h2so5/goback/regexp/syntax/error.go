package syntax

import (
	"regexp/syntax"
	"unicode/utf8"
)

func newErrorBytes(code syntax.ErrorCode, bytes []byte) syntax.Error {
	return syntax.Error{
		Code: code,
		Expr: string(bytes),
	}
}

func newErrorRunes(code syntax.ErrorCode, runes []rune) syntax.Error {
	var rbytes []byte
	for _, r := range runes {
		var lit [utf8.UTFMax]byte
		l := utf8.EncodeRune(lit[:], r)
		rbytes = append(rbytes, lit[:l]...)
	}
	return syntax.Error{
		Code: code,
		Expr: string(rbytes),
	}
}
