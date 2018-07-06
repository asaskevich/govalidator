// Package regexp implements extended regular expression search.
package regexp

// This package provides extended regular expression syntax.
// The implementation does NOT guarantee linear processing time.
//
// If you don't use extended syntax, the golang built-in regex engine will be used transparently.

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/h2so5/goback/regexp/syntax"
)

// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines.
type Regexp interface {

	// Match reports whether the Regexp matches the byte slice b.
	Match(b []byte) bool

	// MatchString reports whether the Regexp matches the string s.
	MatchString(s string) bool

	// Find returns a slice holding the text of the leftmost match in b of the regular expression.
	// A return value of nil indicates no match.
	Find(b []byte) []byte

	// FindIndex returns a two-element slice of integers defining the location of
	// the leftmost match in b of the regular expression.  The match itself is at
	// b[loc[0]:loc[1]].
	// A return value of nil indicates no match.
	FindIndex(b []byte) []int

	// FindSubmatch returns a slice of slices holding the text of the leftmost
	// match of the regular expression in b and the matches, if any, of its
	// subexpressions, as defined by the 'Submatch' descriptions in the package
	// comment.
	// A return value of nil indicates no match.
	FindSubmatch(b []byte) [][]byte

	// FindSubmatchIndex returns a slice holding the index pairs identifying the
	// leftmost match of the regular expression in b and the matches, if any, of
	// its subexpressions, as defined by the 'Submatch' and 'Index' descriptions
	// in the package comment.
	// A return value of nil indicates no match.
	FindSubmatchIndex(b []byte) []int

	// FindString returns a string holding the text of the leftmost match in s of the regular
	// expression.  If there is no match, the return value is an empty string,
	// but it will also be empty if the regular expression successfully matches
	// an empty string.  Use FindStringIndex or FindStringSubmatch if it is
	// necessary to distinguish these cases.
	FindString(s string) string

	// FindStringIndex returns a two-element slice of integers defining the
	// location of the leftmost match in s of the regular expression.  The match
	// itself is at s[loc[0]:loc[1]].
	// A return value of nil indicates no match.
	FindStringIndex(s string) []int

	// FindStringSubmatch returns a slice of strings holding the text of the
	// leftmost match of the regular expression in s and the matches, if any, of
	// its subexpressions, as defined by the 'Submatch' description in the
	// package comment.
	// A return value of nil indicates no match.
	FindStringSubmatch(s string) []string

	// FindStringSubmatchIndex returns a slice holding the index pairs
	// identifying the leftmost match of the regular expression in s and the
	// matches, if any, of its subexpressions, as defined by the 'Submatch' and
	// 'Index' descriptions in the package comment.
	// A return value of nil indicates no match.
	FindStringSubmatchIndex(s string) []int

	// FindAll is the 'All' version of Find; it returns a slice of all successive
	// matches of the expression, as defined by the 'All' description in the
	// package comment.
	// A return value of nil indicates no match.
	FindAll(b []byte, n int) [][]byte

	// FindAllIndex is the 'All' version of FindIndex; it returns a slice of all
	// successive matches of the expression, as defined by the 'All' description
	// in the package comment.
	// A return value of nil indicates no match.
	FindAllIndex(b []byte, n int) [][]int

	// FindAllSubmatchIndex is the 'All' version of FindSubmatchIndex; it returns
	// a slice of all successive matches of the expression, as defined by the
	// 'All' description in the package comment.
	// A return value of nil indicates no match.
	FindAllSubmatchIndex(b []byte, n int) [][]int

	// FindAllSubmatch is the 'All' version of FindSubmatch; it returns a slice
	// of all successive matches of the expression, as defined by the 'All'
	// description in the package comment.
	// A return value of nil indicates no match.
	FindAllSubmatch(b []byte, n int) [][][]byte

	// FindAllString is the 'All' version of FindString; it returns a slice of all
	// successive matches of the expression, as defined by the 'All' description
	// in the package comment.
	// A return value of nil indicates no match.
	FindAllString(s string, n int) []string

	// FindAllStringIndex is the 'All' version of FindStringIndex; it returns a
	// slice of all successive matches of the expression, as defined by the 'All'
	// description in the package comment.
	// A return value of nil indicates no match.
	FindAllStringIndex(s string, n int) [][]int

	// FindAllStringSubmatch is the 'All' version of FindStringSubmatch; it
	// returns a slice of all successive matches of the expression, as defined by
	// the 'All' description in the package comment.
	// A return value of nil indicates no match.
	FindAllStringSubmatch(s string, n int) [][]string

	// FindAllStringSubmatchIndex is the 'All' version of
	// FindStringSubmatchIndex; it returns a slice of all successive matches of
	// the expression, as defined by the 'All' description in the package
	// comment.
	// A return value of nil indicates no match.
	FindAllStringSubmatchIndex(s string, n int) [][]int

	// ReplaceAllFunc returns a copy of src in which all matches of the
	// Regexp have been replaced by the return value of function repl applied
	// to the matched byte slice.  The replacement returned by repl is substituted
	// directly, without using Expand.
	ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte

	// ReplaceAllStringFunc returns a copy of src in which all matches of the
	// Regexp have been replaced by the return value of function repl applied
	// to the matched substring.  The replacement returned by repl is substituted
	// directly, without using Expand.
	ReplaceAllStringFunc(src string, repl func(string) string) string

	// ReplaceAll returns a copy of src, replacing matches of the Regexp
	// with the replacement text repl.  Inside repl, $ signs are interpreted as
	// in Expand, so for instance $1 represents the text of the first submatch.
	ReplaceAll(src, repl []byte) []byte

	// ReplaceAllString returns a copy of src, replacing matches of the Regexp
	// with the replacement string repl.  Inside repl, $ signs are interpreted as
	// in Expand, so for instance $1 represents the text of the first submatch.
	ReplaceAllString(src, repl string) string

	// ReplaceAllLiteral returns a copy of src, replacing matches of the Regexp
	// with the replacement bytes repl.  The replacement repl is substituted directly,
	// without using Expand.
	ReplaceAllLiteral(src, repl []byte) []byte

	// ReplaceAllLiteralString returns a copy of src, replacing matches of the Regexp
	// with the replacement string repl.  The replacement repl is substituted directly,
	// without using Expand.
	ReplaceAllLiteralString(src, repl string) string

	// Expand appends template to dst and returns the result; during the
	// append, Expand replaces variables in the template with corresponding
	// matches drawn from src.  The match slice should have been returned by
	// FindSubmatchIndex.
	//
	// In the template, a variable is denoted by a substring of the form
	// $name or ${name}, where name is a non-empty sequence of letters,
	// digits, and underscores.  A purely numeric name like $1 refers to
	// the submatch with the corresponding index; other names refer to
	// capturing parentheses named with the (?P<name>...) syntax.  A
	// reference to an out of range or unmatched index or a name that is not
	// present in the regular expression is replaced with an empty slice.
	//
	// In the $name form, name is taken to be as long as possible: $1x is
	// equivalent to ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.
	//
	// To insert a literal $ in the output, use $$ in the template.
	Expand(dst []byte, template []byte, src []byte, match []int) []byte

	// ExpandString is like Expand but the template and source are strings.
	// It appends to and returns a byte slice in order to give the calling
	// code control over allocation.
	ExpandString(dst []byte, template string, src string, match []int) []byte

	// Split slices s into substrings separated by the expression and returns a slice of
	// the substrings between those expression matches.
	//
	// The slice returned by this method consists of all the substrings of s
	// not contained in the slice returned by FindAllString. When called on an expression
	// that contains no metacharacters, it is equivalent to strings.SplitN.
	//
	// Example:
	//   s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
	//   // s: ["", "b", "b", "c", "cadaaae"]
	//
	// The count determines the number of substrings to return:
	//   n > 0: at most n substrings; the last substring will be the unsplit remainder.
	//   n == 0: the result is nil (zero substrings)
	//   n < 0: all substrings
	Split(s string, n int) []string

	// LiteralPrefix returns a literal string that must begin any match
	// of the regular expression re.  It returns the boolean true if the
	// literal string comprises the entire regular expression.
	LiteralPrefix() (prefix string, complete bool)

	// NumSubexp returns the number of parenthesized subexpressions in this Regexp.
	NumSubexp() int

	// SubexpNames returns the names of the parenthesized subexpressions
	// in this Regexp.  The name for the first sub-expression is names[1],
	// so that if m is a match slice, the name for m[i] is SubexpNames()[i].
	// Since the Regexp as a whole cannot be named, names[0] is always
	// the empty string.  The slice should not be modified.
	SubexpNames() []string

	// Longest makes future searches prefer the leftmost-longest match.
	// That is, when matching against text, the regexp returns a match that
	// begins as early as possible in the input (leftmost), and among those
	// it chooses a match that is as long as possible.
	Longest()

	// String returns the source text used to compile the regular expression.
	String() string

	// Funcs adds the elements of the argument map to the template's function map.
	Funcs(funcMap syntax.FuncMap)
}

type reg struct {
	*regexp.Regexp
}

func (r *reg) Funcs(funcMap syntax.FuncMap) {}

// Compile parses a regular expression and returns, if successful,
// a Regexp object that can be used to match against text.
func Compile(expr string) (Regexp, error) {
	r, ext, err := syntax.Compile(expr)
	if err != nil {
		return nil, err
	}
	if ext {
		return r, nil
	}
	re, err := regexp.Compile(ignoreComments(expr))
	return &reg{
		Regexp: re,
	}, err
}

// CompileFreeSpacing parses a regular expression like Compile,
// but whitespace characters are ignored and # is parsed as the beggining of a line comment.
func CompileFreeSpacing(expr string) (Regexp, error) {
	r, ext, err := syntax.Compile(ignoreCommentsAndSpaces(expr))
	if err != nil {
		return nil, err
	}
	if ext {
		return r, nil
	}
	re, err := regexp.Compile(ignoreComments(ignoreCommentsAndSpaces(expr)))
	return &reg{
		Regexp: re,
	}, err
}

func compile(expr string) (Regexp, error) {
	r, _, err := syntax.Compile(expr)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// MustCompileFreeSpacing parses a regular expression like MustCompile,
// but whitespace characters are ignored and # is parsed as the beggining of a line comment.
func MustCompileFreeSpacing(str string) Regexp {
	r, err := CompileFreeSpacing(str)
	if err != nil {
		panic(err)
	}
	return r
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompile(str string) Regexp {
	r, err := Compile(str)
	if err != nil {
		panic(err)
	}
	return r
}

func mustCompile(expr string) Regexp {
	r, err := compile(expr)
	if err != nil {
		panic(err)
	}
	return r
}

// Match checks whether a textual regular expression
// matches a byte slice.  More complicated queries need
// to use Compile and the full Regexp interface.
func Match(pattern string, b []byte) (matched bool, err error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.Match(b), nil
}

// MatchString checks whether a textual regular expression
// matches a string.  More complicated queries need
// to use Compile and the full Regexp interface.
func MatchString(pattern string, s string) (matched bool, err error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(s), nil
}

// QuoteMeta returns a string that quotes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text.  For example, QuoteMeta(`[foo]`) returns `\[foo\]`.
func QuoteMeta(s string) string {
	return regexp.QuoteMeta(s)
}

var commentRegexp = regexp.MustCompile(`(^|[^\\]|\\\\)(\(\?#[^)]*\))`)

func ignoreComments(expr string) string {
	return commentRegexp.ReplaceAllString(expr, "$1")
}

var lineCommentRegexp = regexp.MustCompile(`(?m)(?:\\#)|(\#)|(?:#.*$)`)
var spaceRegexp = regexp.MustCompile(`( )|(?:)`)

func ignoreCommentsAndSpaces(expr string) string {
	var res string
loop:
	for _, line := range strings.Split(expr, "\n") {
		meta := false
		charClass := false
		for _, c := range line {
			if meta {
				charClass = false
				switch c {
				case '\\':
					res += "\\"
				case '#':
					res += "#"
				case ' ':
					res += " "
				default:
					res += fmt.Sprintf("\\%c", c)
				}
				meta = false
			} else {
				if charClass {
					if c == ' ' {
						res += fmt.Sprintf("%c", c)
					}
					charClass = false
				} else {
					if c == '[' {
						charClass = true
					}
				}
				switch c {
				case '\\':
					meta = true
				case ' ':
				case '\v':
				case '\t':
				case '\f':
				case '\r':
				case '[':
					res += "["
				case '#':
					continue loop
				default:
					res += fmt.Sprintf("%c", c)
				}
			}
		}
	}
	return res
}
