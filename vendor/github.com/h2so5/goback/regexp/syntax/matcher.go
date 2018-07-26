package syntax

import (
	"regexp/syntax"
	"unicode"
)

type charNodeMatcher interface {
	Match(rune, syntax.Flags) bool
}

type reverseMatcher struct {
	M charNodeMatcher
}

func (m reverseMatcher) Match(r rune, flags syntax.Flags) bool {
	return !m.M.Match(r, flags)
}

type unicodeMatcher struct {
	R *unicode.RangeTable
}

func (m unicodeMatcher) Match(r rune, flags syntax.Flags) bool {
	return unicode.Is(m.R, r)
}

type digitsMatcher struct{}

func (m digitsMatcher) Match(r rune, flags syntax.Flags) bool {
	return '0' <= r && r <= '9'
}

type whitespaceMatcher struct{}

func (m whitespaceMatcher) Match(r rune, flags syntax.Flags) bool {
	switch r {
	case '\t', '\n', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}

func isASCIIWord(r rune) bool {
	return ('0' <= r && r <= '9') ||
		('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z') || r == '_'
}

type wordMatcher struct{}

func (m wordMatcher) Match(r rune, flags syntax.Flags) bool {
	return isASCIIWord(r)
}

type alphanumericMatcher struct {
}

func (m alphanumericMatcher) Match(r rune, flags syntax.Flags) bool {
	return ('0' <= r && r <= '9') ||
		('a' <= r && r <= 'z') ||
		('A' <= r && r <= 'Z')
}

type alphabeticMatcher struct {
}

func (m alphabeticMatcher) Match(r rune, flags syntax.Flags) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}

type asciiMatcher struct {
}

func (m asciiMatcher) Match(r rune, flags syntax.Flags) bool {
	return 0 <= r && r <= 0x7F
}

type blankMatcher struct {
}

func (m blankMatcher) Match(r rune, flags syntax.Flags) bool {
	return r == '\t' || r == ' '
}

type controlMatcher struct {
}

func (m controlMatcher) Match(r rune, flags syntax.Flags) bool {
	return (0 <= r && r <= 0x1F) || r == 0x7F
}

type graphicalMatcher struct {
}

func (m graphicalMatcher) Match(r rune, flags syntax.Flags) bool {
	return '!' <= r && r <= '~'
}

type lowerMatcher struct {
}

func (m lowerMatcher) Match(r rune, flags syntax.Flags) bool {
	return 'a' <= r && r <= 'z'
}

type printableMatcher struct {
}

func (m printableMatcher) Match(r rune, flags syntax.Flags) bool {
	return ' ' <= r && r <= '~'
}

func isASCIIPunct(r rune) bool {
	return ('!' <= r && r <= '/') || (':' <= r && r <= '@') ||
		('[' <= r && r <= '`') || ('{' <= r && r <= '~')
}

type punctuationMatcher struct {
}

func (m punctuationMatcher) Match(r rune, flags syntax.Flags) bool {
	return isASCIIPunct(r)
}

type upperMatcher struct {
}

func (m upperMatcher) Match(r rune, flags syntax.Flags) bool {
	return 'A' <= r && r <= 'Z'
}

func isASCIIXdigit(r rune) bool {
	return ('0' <= r && r <= '9') ||
		('a' <= r && r <= 'f') ||
		('A' <= r && r <= 'F')
}

type xdigitMatcher struct {
}

func (m xdigitMatcher) Match(r rune, flags syntax.Flags) bool {
	return isASCIIXdigit(r)
}
