package syntax

import (
	"bytes"
	"regexp/syntax"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type regexp struct {
	root        node
	expr        string
	subexpNames []string
	subexpMap   map[string]int
	longest     bool
	funcs       []FuncMap
}

func (re *regexp) NumSubexp() int {
	return len(re.subexpNames) - 1
}

func (re *regexp) Match(b []byte) bool {
	return len(re.Find(b)) > 0
}

func (re *regexp) MatchString(s string) bool {
	return re.Match([]byte(s))
}

func (re *regexp) Find(b []byte) []byte {
	loc := re.FindIndex(b)
	if len(loc) == 0 {
		return nil
	}
	return b[loc[0]:loc[1]]
}

func (re *regexp) FindIndex(b []byte) []int {
	p, comp := re.literalPrefix()
	i := bytes.Index(b, p)
	if i < 0 {
		return nil
	} else if comp {
		return []int{i, i + len(p)}
	}
	loc := re.FindSubmatchIndex(b)
	if len(loc) == 0 {
		return nil
	}
	return loc[:2]
}

func (re *regexp) FindSubmatch(b []byte) [][]byte {
	var ret [][]byte
	loc := re.FindSubmatchIndex(b)
	for i := 0; i < len(loc)/2; i++ {
		ret = append(ret, b[loc[i*2]:loc[i*2+1]])
	}
	return ret
}

func (re *regexp) FindSubmatchIndex(b []byte) []int {
	return re.findSubmatchIndex(b, 0)
}

func (re *regexp) findSubmatchIndex(b []byte, f int) []int {
	offset := f

	fixed := false
	if b, _ := re.root.Hint()[hintFixedBeginning].(bool); b {
		fixed = true
	}

	p, comp := re.literalPrefix()
	i := bytes.Index(b[offset:], p)
	if i < 0 {
		return nil
	} else {
		offset += i
		if comp && re.NumSubexp() == 0 {
			return []int{offset, offset + len(p)}
		}
	}

	for {
		f := re.root.Fiber(input{
			b: b[offset:],
			o: b, begin: offset,
			funcs: re.funcs,
		})
		o, err := f.Resume()
		if err == nil {
			if re.longest {
				for {
					a, err := f.Resume()
					if err != nil {
						break
					} else if a.offset > o.offset {
						o = a
					}
				}
			}
			loc := make([]int, 0, re.NumSubexp()*2)
			loc = append(loc, []int{offset, offset + o.offset}...)
			for i := 1; i <= re.NumSubexp(); i++ {
				if sub, ok := o.sub.i[i]; ok {
					loc = append(loc, sub.begin, sub.begin+len(sub.b))
				} else {
					loc = append(loc, -1, -1)
				}
			}
			return loc
		}
		if fixed || len(b[offset:]) == 0 {
			break
		}
		_, s := utf8.DecodeRune(b[offset:])
		offset += s
	}
	return nil
}

func (re *regexp) FindAllString(s string, n int) []string {
	var ret []string
	for _, b := range re.FindAll([]byte(s), n) {
		ret = append(ret, string(b))
	}
	return ret
}

func (re *regexp) FindAllStringIndex(s string, n int) [][]int {
	return re.FindAllIndex([]byte(s), n)
}

func (re *regexp) FindAll(b []byte, n int) [][]byte {
	var ret [][]byte
	for _, loc := range re.FindAllIndex(b, n) {
		ret = append(ret, b[loc[0]:loc[1]])
	}
	return ret
}

func (re *regexp) FindAllIndex(b []byte, n int) [][]int {
	var ret [][]int
	for _, loc := range re.FindAllSubmatchIndex(b, n) {
		ret = append(ret, loc[:2])
	}
	return ret
}

func (re *regexp) FindAllStringSubmatchIndex(s string, n int) [][]int {
	return re.FindAllSubmatchIndex([]byte(s), n)
}

func (re *regexp) FindAllStringSubmatch(s string, n int) [][]string {
	var ret [][]string
	for _, m := range re.FindAllSubmatch([]byte(s), n) {
		var sub []string
		for _, b := range m {
			sub = append(sub, string(b))
		}
		ret = append(ret, sub)
	}
	return ret
}

func (re *regexp) FindAllSubmatch(b []byte, n int) [][][]byte {
	var ret [][][]byte
	for _, m := range re.FindAllSubmatchIndex(b, n) {
		var sub [][]byte
		for i := 0; i < len(m)/2; i++ {
			sub = append(sub, b[m[i*2]:m[i*2+1]])
		}
		ret = append(ret, sub)
	}
	return ret
}

func (re *regexp) FindAllSubmatchIndex(b []byte, n int) [][]int {
	var ret [][]int
	offset := 0
	for i := 0; i < n || n < 0; i++ {
		m := re.findSubmatchIndex(b, offset)
		if len(m) == 0 {
			break
		}
		last := len(ret) - 1
		if m[0] == m[1] && last >= 0 {
			if m[0] != ret[last][1] {
				ret = append(ret, m)
			}
		} else {
			ret = append(ret, m)
		}
		if len(b[offset:]) == 0 {
			break
		}
		if m[1] > offset {
			offset = m[1]
		} else {
			_, s := utf8.DecodeRune(b[offset:])
			offset += s
		}
	}
	return ret
}

func (re *regexp) FindString(s string) string {
	return string(re.Find([]byte(s)))
}

func (re *regexp) FindStringIndex(s string) []int {
	return re.FindIndex([]byte(s))
}

func (re *regexp) FindStringSubmatch(s string) []string {
	var ret []string
	for _, b := range re.FindSubmatch([]byte(s)) {
		ret = append(ret, string(b))
	}
	return ret
}

func (re *regexp) FindStringSubmatchIndex(s string) []int {
	return re.FindSubmatchIndex([]byte(s))
}

func (re *regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte {
	sub, sep, _ := re.split(src)
	if len(sep) == 0 {
		return append([]byte(nil), src...)
	}
	var ret []byte
	for i, s := range sub[:len(sub)-1] {
		ret = append(append(ret, s...), repl(sep[i])...)
	}
	ret = append(ret, sub[len(sub)-1]...)
	return ret
}

func (re *regexp) ReplaceAllStringFunc(src string, repl func(string) string) string {
	return string(re.ReplaceAllFunc([]byte(src), func(b []byte) []byte {
		return []byte(repl(string(b)))
	}))
}

func (re *regexp) ReplaceAllLiteral(src, repl []byte) []byte {
	return re.ReplaceAllFunc(src, func([]byte) []byte {
		return repl
	})
}

func (re *regexp) ReplaceAllLiteralString(src, repl string) string {
	return string(re.ReplaceAllLiteral([]byte(src), []byte(repl)))
}

func (re *regexp) ReplaceAll(src, repl []byte) []byte {
	sub, sep, match := re.split(src)
	if len(sep) == 0 {
		return append([]byte(nil), src...)
	}
	var ret []byte
	for i, s := range sub[:len(sub)-1] {
		ret = append(append(ret, s...), re.Expand(nil, repl, src, match[i])...)
	}
	ret = append(ret, sub[len(sub)-1]...)
	return ret
}

func (re *regexp) ReplaceAllString(src, repl string) string {
	return string(re.ReplaceAll([]byte(src), []byte(repl)))
}

func (re *regexp) split(b []byte) ([][]byte, [][]byte, [][]int) {
	var idx [][]int
	var sep [][]byte
	var match [][]int
	before := 0
	for _, i := range re.FindAllSubmatchIndex(b, -1) {
		idx = append(idx, []int{before, i[0]})
		sep = append(sep, b[i[0]:i[1]])
		match = append(match, i)
		before = i[1]
	}
	idx = append(idx, []int{before, len(b)})
	var sub [][]byte
	for _, i := range idx {
		sub = append(sub, b[i[0]:i[1]])
	}
	return sub, sep, match
}

// From http://golang.org/src/regexp/regexp.go
func (re *regexp) Split(s string, n int) []string {

	if n == 0 {
		return nil
	}

	if len(re.expr) > 0 && len(s) == 0 {
		return []string{""}
	}

	matches := re.FindAllStringIndex(s, n)
	strings := make([]string, 0, len(matches))

	beg := 0
	end := 0
	for _, match := range matches {
		if n > 0 && len(strings) >= n-1 {
			break
		}

		end = match[0]
		if match[1] != 0 {
			strings = append(strings, s[beg:end])
		}
		beg = match[1]
	}

	if end != len(s) {
		strings = append(strings, s[beg:])
	}

	return strings
}

func (re *regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte {
	var res []byte
	meta := false
	runes := bytes.Runes(template)
	for i := 0; i < len(runes); i++ {
		switch {
		case runes[i] == '$':
			if meta {
				res = append(res, '$')
			}
			meta = !meta
		default:
			if meta {
				meta = false
				name, l := re.parseTemplate(runes[i:])
				if l > 0 {
					idx, err := strconv.Atoi(name)
					if err == nil && strconv.Itoa(idx) == name {
						if idx < len(match)/2 {
							res = append(res, src[match[idx*2]:match[idx*2+1]]...)
						}
					} else {
						if idx, ok := re.subexpMap[name]; ok {
							res = append(res, src[match[idx*2]:match[idx*2+1]]...)
						}
					}
					i += l - 1
					continue
				}
			}
			var lit [utf8.UTFMax]byte
			l := utf8.EncodeRune(lit[:], runes[i])
			res = append(res, lit[:l]...)
		}
	}
	copy(dst, res)
	return res
}

func (re *regexp) ExpandString(dst []byte, template string, src string, match []int) []byte {
	return re.Expand(dst, []byte(template), []byte(src), match)
}

func (re *regexp) parseTemplate(exp []rune) (string, int) {
	if len(exp) == 0 {
		return "", 0
	}
	if exp[0] == '{' {
		for i, r := range exp[1:] {
			if r == '}' {
				name, l := re.parseTemplate(exp[1 : i+1])
				return name, l + 2
			}
		}
	} else {
		var name string
		for i, r := range exp {
			if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
				var lit [utf8.UTFMax]byte
				l := utf8.EncodeRune(lit[:], r)
				name += string(lit[:l])
			} else {
				return name, i
			}
		}
		return name, len(exp)
	}
	return "", 0
}

func (re *regexp) SubexpNames() []string {
	return re.subexpNames
}

func (re *regexp) LiteralPrefix() (prefix string, complete bool) {
	b, comp := re.literalPrefix()
	return string(b), comp
}

func (re *regexp) literalPrefix() (prefix []byte, complete bool) {
	return re.root.LiteralPrefix()
}

func (re *regexp) Longest() {
	re.longest = true
}

func (re *regexp) String() string {
	return re.expr
}

// Context represents a matching context used in a inline function
type Context struct {
	Data    []byte
	Cursor  int
	Matches map[interface{}][]int
}

// FuncMap is the type of the map defining the mapping from names to functions.
type FuncMap map[string]func(ctx Context) interface{}

func (re *regexp) Funcs(funcMap FuncMap) {
	re.funcs = append(re.funcs, funcMap)
}

// Compile parses a regular expression and returns, if successful,
// a Regexp object that can be used to match against text.
func Compile(expr string) (re *regexp, extended bool, err error) {
	p := parser{}
	flags := syntax.OneLine | syntax.PerlX
	n, subexp, err := p.parse([]byte(expr), flags)
	if err != nil {
		return nil, false, err
	}
	m := make(map[string]int)
	for i, n := range subexp {
		if len(n) > 0 {
			m[n] = i
		}
	}
	return &regexp{
		root:        n,
		expr:        expr,
		subexpNames: subexp,
		subexpMap:   m,
	}, n.IsExtended(), nil
}
