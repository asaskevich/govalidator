package syntax

import (
	"fmt"
	"reflect"
	"regexp/syntax"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type parser struct {
	groupIndex  int
	subexpNames []string
}

func (p *parser) parse(reg []byte, flags syntax.Flags) (n node, subexp []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(syntax.Error); ok {
				err = fmt.Errorf("error parsing regexp: %s", e.Error())
			} else {
				panic(r)
			}
		}
	}()

	p.groupIndex = 0
	p.subexpNames = nil

	runes := make([]rune, 0, len(reg))
	b := reg
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		if r == utf8.RuneError {
			panic(newErrorBytes(syntax.ErrInvalidUTF8, reg))
		}
		runes = append(runes, r)
		b = b[size:]
	}

	p.groupIndex = -1
	return p.group(runes, flags), p.subexpNames, nil
}

func parseBackref(exp []rune) (string, int) {
	if len(exp) == 0 {
		return "", 0
	}
	if exp[0] == '{' {
		for i, r := range exp[1:] {
			if r == '}' {
				name, l := parseBackref(exp[1 : i+1])
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

func applyFlags(f syntax.Flags, m map[syntax.Flags]int) syntax.Flags {
	for k, v := range m {
		if v == 1 {
			f |= k
		} else if v == -1 {
			f &= ^k
		}
	}
	return f
}

const (
	wrapperNone = iota
	wrapperLookahead
	wrapperNegativeLookahead
	wrapperLookbehind
	wrapperNegativeLookbehind
)

func (p *parser) group(runes []rune, flags syntax.Flags) node {
	g := groupNode{
		N: []node{},
	}

	mflags := map[syntax.Flags]int{}
	meta := false
	r := runes
	indexed := true
	wrapper := wrapperNone

	exp := append(append([]rune{'('}, runes...), ')')
	if len(r) >= 2 && r[0] == '?' {
		switch {
		case r[1] == '>':
			g.Atomic = true
			indexed = false
			r = r[2:]
		case r[1] == ':':
			indexed = false
			r = r[2:]
		case r[1] == '#':
			return g
		case r[1] == '=':
			indexed = false
			wrapper = wrapperLookahead
			r = r[2:]
		case r[1] == '{':
			if len(runes) <= 3 || runes[len(runes)-1] != '}' {
				panic(newErrorRunes(syntax.ErrInvalidPerlOp, exp))
			}
			var name []byte
			for _, r := range runes[2 : len(runes)-1] {
				var lit [utf8.UTFMax]byte
				l := utf8.EncodeRune(lit[:], r)
				name = append(name, lit[:l]...)
			}
			return funcNode{Name: strings.TrimSpace(string(name))}

		case len(r) >= 3 && r[1] == '<' && r[2] == '=':
			indexed = false
			wrapper = wrapperLookbehind
			r = r[3:]
		case len(r) >= 3 && r[1] == '<' && r[2] == '!':
			indexed = false
			wrapper = wrapperNegativeLookbehind
			r = r[3:]
		case r[1] == '!':
			indexed = false
			wrapper = wrapperNegativeLookahead
			r = r[2:]
		case r[1] == 'P':
			if len(r) >= 3 && r[2] == '<' {
				r = r[3:]
				var name []rune
				for i, e := range r {
					if e == '>' {
						name = r[:i]
						break
					} else if e != '_' &&
						!('a' <= e && e <= 'z') &&
						!('A' <= e && e <= 'Z') &&
						!('0' <= e && e <= '9') {
						break
					}
				}
				if len(name) == 0 {
					panic(newErrorRunes(syntax.ErrInvalidNamedCapture, exp))
				}

				r = r[len(name)+1:]

				var rbytes []byte
				for _, r := range name {
					var lit [utf8.UTFMax]byte
					l := utf8.EncodeRune(lit[:], r)
					rbytes = append(rbytes, lit[:l]...)
				}
				g.Name = string(rbytes)

			} else {
				panic(newErrorRunes(syntax.ErrInvalidPerlOp, exp))
			}
		default:
			f := 1
			indexed = false
			internal := false
		loop:
			for r = r[1:]; len(r) > 0; r = r[1:] {
				switch r[0] {
				case 'i':
					mflags[syntax.FoldCase] = f
				case 'm':
					mflags[syntax.OneLine] = -f
				case 's':
					mflags[syntax.DotNL] = f
				case 'U':
					mflags[syntax.NonGreedy] = f
				case ':':
					internal = true
					r = r[1:]
					break loop
				case '-':
					if f == 1 {
						f = -1
					} else {
						panic(newErrorRunes(syntax.ErrInvalidPerlOp, exp))
					}
				default:
					panic(newErrorRunes(syntax.ErrInvalidPerlOp, exp))
				}
			}
			if !internal {
				return flagNode{Flags: mflags}
			}
		}
	}

	flags = applyFlags(flags, mflags)

	if indexed {
		p.groupIndex++
		g.Index = p.groupIndex
		p.subexpNames = append(p.subexpNames, g.Name)
	}

	for len(r) > 0 {
		rx := runes[:len(runes)-len(r)]
		if meta {
			meta = false
			switch {
			case r[0] == '0':
				n, _ := p.fetchLiteral([]rune{rune(0)}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case '1' <= r[0] && r[0] <= '7':
				size := 1
				oct := []int{int(r[0] - '0'), 0, 0}
				if len(r) >= 2 && '0' <= r[1] && r[1] <= '7' {
					oct = append([]int{int(r[1] - '0')}, oct...)
					size++
				}
				if len(r) >= 3 && '0' <= r[2] && r[2] <= '7' {
					oct = append([]int{int(r[2] - '0')}, oct...)
					size++
				}
				i := oct[0] + oct[1]*8 + oct[2]*64
				n, _ := p.fetchLiteral([]rune{rune(i)}, flags)
				r = r[size:]
				g.N = append(g.N, n)
			case r[0] == 'x':
				n, size := p.fetchHexCode(r, flags)
				r = r[size:]
				g.N = append(g.N, n)
			case r[0] == 'k':
				name, size := parseBackref(r[1:])
				if size > 0 {
					n := backRefNode{Flags: flags}
					idx, err := strconv.Atoi(name)
					if err == nil && strconv.Itoa(idx) == name {
						n.Index = idx
					} else {
						n.Name = name
					}
					r = r[size+1:]
					g.N = append(g.N, n)
				} else {
					panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, r[0])))
				}
			case r[0] == 'd':
				n := charNode{
					Flags: flags,
					Matcher: []charNodeMatcher{
						digitsMatcher{},
					},
				}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'D':
				n := charNode{
					Flags: flags,
					Matcher: []charNodeMatcher{
						reverseMatcher{M: digitsMatcher{}},
					},
				}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 's':
				n := charNode{
					Flags: flags,
					Matcher: []charNodeMatcher{
						whitespaceMatcher{},
					},
				}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'S':
				n := charNode{
					Flags: flags,
					Matcher: []charNodeMatcher{
						reverseMatcher{M: whitespaceMatcher{}},
					},
				}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'w':
				n := charNode{
					Flags: flags,
					Matcher: []charNodeMatcher{
						wordMatcher{},
					},
				}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'W':
				n := charNode{
					Flags: flags,
					Matcher: []charNodeMatcher{
						reverseMatcher{M: wordMatcher{}},
					},
				}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'p' || r[0] == 'P':
				m, size := p.fetchUnicodeClass(r)
				n := charNode{
					Flags:   flags,
					Matcher: []charNodeMatcher{m},
				}
				r = r[size:]
				g.N = append(g.N, n)
			case r[0] == 'A':
				n := beginNode{Flags: flags}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'z':
				n := endNode{Flags: flags}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'a':
				n, _ := p.fetchLiteral([]rune{'\a'}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'f':
				n, _ := p.fetchLiteral([]rune{'\f'}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 't':
				n, _ := p.fetchLiteral([]rune{'\t'}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'n':
				n, _ := p.fetchLiteral([]rune{'\n'}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'r':
				n, _ := p.fetchLiteral([]rune{'\r'}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'v':
				n, _ := p.fetchLiteral([]rune{'\v'}, flags)
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'Q':
				l := 1
				for i := 1; i < len(r); i++ {
					if i+1 < len(r) && r[i] == '\\' && r[i+1] == 'E' {
						l += 2
						break
					}
					n, _ := p.fetchLiteral(r[i:], flags)
					g.N = append(g.N, n)
					l++
				}
				r = r[l:]
			case r[0] == 'E':
				panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, r[0])))
			case r[0] == 'b':
				n := wordBoundaryNode{}
				r = r[1:]
				g.N = append(g.N, n)
			case r[0] == 'B':
				n := wordBoundaryNode{Reversed: true}
				r = r[1:]
				g.N = append(g.N, n)
			case isASCIIPunct(r[0]):
				n, size := p.fetchLiteral(r, flags)
				r = r[size:]
				g.N = append(g.N, n)
			default:
				panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, r[0])))
			}
		} else {
			switch r[0] {
			case '\\':
				meta = true
				r = r[1:]
			case '.':
				n := anyCharNode{
					Flags: flags,
				}
				r = r[1:]
				g.N = append(g.N, n)
			case '^':
				n := beginNode{Flags: flags, Line: true}
				r = r[1:]
				g.N = append(g.N, n)
			case '$':
				n := endNode{Flags: flags, Line: true}
				r = r[1:]
				g.N = append(g.N, n)
			case '?', '*', '+': // repeat
				n, size := p.fetchRepeat(r, flags)
				r = r[size:]
				g.N = append(g.N, n)
			case '|':
				r = r[1:]
				g.N = append(g.N, alterNode{})
			case '[':
				n, size := p.fetchCharClass(r, flags)
				r = r[size:]
				g.N = append(g.N, n)
			case '{':
				n, size := p.fetchRange(r)
				r = r[size:]
				g.N = append(g.N, n)
			case '(': // group
				n, size := p.fetchGroup(r, flags)
				if n == nil {
					panic(newErrorRunes(syntax.ErrMissingParen, append(rx, r...)))
				}
				r = r[size:]
				if fn, ok := n.(flagNode); ok {
					flags = applyFlags(flags, fn.Flags)
				} else {
					g.N = append(g.N, n)
				}
			case ')':
				panic(newErrorRunes(syntax.ErrUnexpectedParen, append(rx, r...)))
			default:
				n, size := p.fetchLiteral(r, flags)
				r = r[size:]
				g.N = append(g.N, n)
			}
		}
	}

	if meta {
		panic(newErrorRunes(syntax.ErrTrailingBackslash, runes))
	}

	g.N = p.concatRepetitions(g.N)
	g.N = p.concatLiterals(g.N)
	g.N = p.concatAlternations(g.N)
	g.N = p.removeSequentialBoundaries(g.N)

	_, max := g.MinMax()

	switch wrapper {
	case wrapperLookahead:
		return lookaheadNode{N: g}
	case wrapperNegativeLookahead:
		return lookaheadNode{N: g, Negative: true}
	case wrapperLookbehind:
		if max < 0 {
			panic(newErrorRunes(syntax.ErrInvalidPerlOp, exp))
		}
		return lookbehindNode{N: g}
	case wrapperNegativeLookbehind:
		if max < 0 {
			panic(newErrorRunes(syntax.ErrInvalidPerlOp, exp))
		}
		return lookbehindNode{N: g, Negative: true}
	}
	return g
}

func (p *parser) fetchLiteral(runes []rune, flags syntax.Flags) (node, int) {
	var lit [utf8.UTFMax]byte
	l := utf8.EncodeRune(lit[:], runes[0])
	return literalNode{L: lit[:l], Flags: flags}, 1
}

func (p *parser) fetchUnicodeClass(runes []rune) (charNodeMatcher, int) {
	var m charNodeMatcher
	size := 0
	if len(runes) < 2 {
		panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes[0])))
	}
	if runes[1] == '{' {
		var name []byte
		for i, r := range runes[2:] {
			if r == '}' {
				size = i + 3
			} else {
				var lit [utf8.UTFMax]byte
				l := utf8.EncodeRune(lit[:], r)
				name = append(name, lit[:l]...)
			}
		}
		if size == 0 {
			panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes...)))
		}
		if c, ok := unicode.Scripts[string(name)]; ok {
			m = &unicodeMatcher{R: c}
		} else {
			panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes[:2]...)))
		}
	} else {
		var lit [utf8.UTFMax]byte
		l := utf8.EncodeRune(lit[:], runes[1])
		if c, ok := unicode.Categories[string(lit[:l])]; ok {
			m = &unicodeMatcher{R: c}
			size = 2
		} else {
			panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes[:2]...)))
		}
	}
	if runes[0] == 'P' {
		m = &reverseMatcher{M: m}
	}
	return m, size
}

func (p *parser) fetchHexCode(runes []rune, flags syntax.Flags) (node, int) {
	var hex []byte
	size := 0
	if len(runes) < 3 {
		panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes...)))
	}
	if runes[1] == '{' {
		for i, r := range runes[2:] {
			if r == '}' {
				size = i + 3
			} else if isASCIIXdigit(r) {
				hex = append(hex, byte(r))
			} else {
				panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes[:i+1]...)))
			}
		}
		if size == 0 {
			panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes...)))
		}
	} else {
		if isASCIIXdigit(runes[1]) && isASCIIXdigit(runes[2]) {
			hex = []byte{byte(runes[1]), byte(runes[2])}
			size = 3
		} else {
			panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes[:2]...)))
		}
	}
	i := 0
	fmt.Sscanf(string(hex), "%x", &i)
	if rune(i) > unicode.MaxRune {
		panic(newErrorRunes(syntax.ErrInvalidEscape, append([]rune{'\\'}, runes...)))
	}
	n, _ := p.fetchLiteral([]rune{rune(i)}, flags)
	return n, size
}

func runesToInt(runes []rune) int {
	i := 0
	b := 1
	for n := range runes {
		i += (int(runes[len(runes)-n-1]) - '0') * b
		b *= 10
	}
	return i
}

func (p *parser) fetchCharClass(runes []rune, flags syntax.Flags) (node, int) {
	l := 1
	r := runes[1:]
	if len(r) < 2 {
		panic(newErrorRunes(syntax.ErrMissingBracket, runes))
	}

	reversed := false
	if r[0] == '^' {
		reversed = true
		r = r[1:]
		l++
	}

	var exp []rune
	meta := false
	depth := 1
	for i, e := range r {
		l++
		if meta {
			meta = false
		} else if e == '\\' {
			meta = true
		} else if e == '[' {
			depth++
		} else if i > 0 && e == ']' {
			depth--
			exp = r[:i]
			if depth == 0 {
				break
			}
		}
	}

	if len(exp) == 0 {
		panic(newErrorRunes(syntax.ErrMissingBracket, runes))
	}

	n := charNode{
		Flags:    flags,
		Matcher:  p.parseCharClass(exp),
		Reversed: reversed,
	}
	return n, l
}

type rangeMatcher struct {
	B, E rune
}

func (m rangeMatcher) Match(r rune, flags syntax.Flags) bool {
	if m.B <= r && r <= m.E {
		return true
	}
	return false
}

type mapMatcher struct {
	M map[rune]int
}

func (m mapMatcher) Match(r rune, flags syntax.Flags) bool {
	if _, ok := m.M[r]; ok {
		return true
	}
	return false
}

func (p *parser) parseCharClass(exp []rune) []charNodeMatcher {
	var m []charNodeMatcher
	var ranges []rangeMatcher
	runes := map[rune]int{}

	meta := false
	for r := exp; len(r) > 0; r = r[1:] {
		if meta {
			switch r[0] {
			case 'd':
				m = append(m, digitsMatcher{})
			case 'D':
				m = append(m, reverseMatcher{M: digitsMatcher{}})
			case 's':
				m = append(m, whitespaceMatcher{})
			case 'S':
				m = append(m, reverseMatcher{M: whitespaceMatcher{}})
			case 'w':
				m = append(m, wordMatcher{})
			case 'W':
				m = append(m, reverseMatcher{M: wordMatcher{}})
			case 'a':
				runes['\a'] = 0
			case 'f':
				runes['\f'] = 0
			case 't':
				runes['\t'] = 0
			case 'n':
				runes['\n'] = 0
			case 'r':
				runes['\r'] = 0
			case 'v':
				runes['\v'] = 0
			case 'p', 'P':
				u, size := p.fetchUnicodeClass(r)
				m = append(m, u)
				r = r[size-1:]
			default:
				runes[r[0]] = 0
			}
			meta = false
		} else if r[0] == '\\' {
			meta = true
		} else if len(r) > 3 && r[0] == '[' && r[1] == ':' {
			var name []byte
			offset := 0
			for i := 2; i < len(r)-1; i++ {
				if r[i] == ':' && r[i+1] == ']' {
					offset = i + 1
					break
				}
				var lit [utf8.UTFMax]byte
				l := utf8.EncodeRune(lit[:], r[i])
				name = append(name, lit[:l]...)
			}
			if offset > 0 {
				switch string(name) {
				case "alnum":
					m = append(m, alphanumericMatcher{})
				case "^alnum":
					m = append(m, reverseMatcher{M: alphanumericMatcher{}})
				case "alpha":
					m = append(m, alphabeticMatcher{})
				case "^alpha":
					m = append(m, reverseMatcher{M: alphabeticMatcher{}})
				case "ascii":
					m = append(m, asciiMatcher{})
				case "^ascii":
					m = append(m, reverseMatcher{M: asciiMatcher{}})
				case "blank":
					m = append(m, blankMatcher{})
				case "^blank":
					m = append(m, reverseMatcher{M: blankMatcher{}})
				case "cntrl":
					m = append(m, controlMatcher{})
				case "^cntrl":
					m = append(m, reverseMatcher{M: controlMatcher{}})
				case "graph":
					m = append(m, graphicalMatcher{})
				case "^graph":
					m = append(m, reverseMatcher{M: graphicalMatcher{}})
				case "lower":
					m = append(m, lowerMatcher{})
				case "^lower":
					m = append(m, reverseMatcher{M: lowerMatcher{}})
				case "print":
					m = append(m, printableMatcher{})
				case "^print":
					m = append(m, reverseMatcher{M: printableMatcher{}})
				case "punct":
					m = append(m, punctuationMatcher{})
				case "^punct":
					m = append(m, reverseMatcher{M: punctuationMatcher{}})
				case "upper":
					m = append(m, upperMatcher{})
				case "^upper":
					m = append(m, reverseMatcher{M: upperMatcher{}})
				case "xdigit":
					m = append(m, xdigitMatcher{})
				case "^xdigit":
					m = append(m, reverseMatcher{M: xdigitMatcher{}})
				case "digit":
					m = append(m, digitsMatcher{})
				case "^digit":
					m = append(m, reverseMatcher{M: digitsMatcher{}})
				case "word":
					m = append(m, wordMatcher{})
				case "^word":
					m = append(m, reverseMatcher{M: wordMatcher{}})
				case "space":
					m = append(m, whitespaceMatcher{})
				case "^space":
					m = append(m, reverseMatcher{M: whitespaceMatcher{}})
				default:
					panic(newErrorRunes(syntax.ErrInvalidCharRange, r[:offset+1]))
				}
				r = r[offset:]
			} else {
				runes[r[0]] = 0
			}
		} else if len(r) >= 3 && r[1] == '-' {
			if r[0] > r[2] {
				panic(newErrorRunes(syntax.ErrInvalidCharRange, r[:3]))
			}
			ranges = append(ranges, rangeMatcher{B: r[0], E: r[2]})
			r = r[2:]
		} else {
			runes[r[0]] = 0
		}
	}

	for _, r := range ranges {
		m = append(m, r)
	}

	// TODO: merge ranges
	for r := range runes {
		for _, g := range ranges {
			if g.B <= r && r <= g.E {
				delete(runes, r)
				break
			}
		}
	}

	if len(runes) > 0 {
		m = append(m, mapMatcher{M: runes})
	}
	return m
}

func (p *parser) fetchRange(runes []rune) (node, int) {
	lit := literalNode{L: []byte{'{'}}
	exp := runes[1:]
	if len(exp) == 0 {
		return lit, 1
	}

	var w [2][]rune
	index := 0
	for i, r := range exp {
		switch {
		case r == '0':
			w[index] = append(w[index], r)
		case '1' <= r && r <= '9':
			if len(w[index]) > 0 && w[index][0] == '0' {
				return lit, 1
			}
			w[index] = append(w[index], r)
		case r == '}':
			if len(w[0]) == 0 && len(w[1]) == 0 {
				return lit, 1
			}
			b := 0
			e := -1
			if len(w[0]) > 0 {
				b = runesToInt(w[0])
			}
			if len(w[1]) > 0 {
				e = runesToInt(w[1])
			}
			if index == 0 {
				e = b
			}
			l := i + 2
			reluctant := false
			atomic := false
			if len(exp) > i+1 {
				if exp[i+1] == '?' {
					reluctant = true
					l++
				} else if exp[i+1] == '+' {
					atomic = true
					l++
				}
			}
			if (e >= 0 && e < b) || b > 1000 || e > 1000 {
				panic(newErrorRunes(syntax.ErrInvalidRepeatSize, runes[:l]))
			}
			return repeatNode{
				Min: b, Max: e,
				Reluctant: reluctant,
				Atomic:    atomic,
				Exp:       runes[:l],
			}, l
		case r == ',':
			if index == 0 && len(w[index]) > 0 {
				index = 1
			} else {
				return lit, 1
			}
		default:
			return lit, 1
		}
	}
	return lit, 1
}

func (p *parser) fetchRepeat(runes []rune, flags syntax.Flags) (node, int) {
	reluctant := false
	atomic := false
	l := 1
	if len(runes) >= 2 {
		if runes[1] == '?' {
			reluctant = true
			l++
		} else if runes[1] == '+' {
			atomic = true
			l++
		}
	}
	if flags&syntax.NonGreedy != 0 {
		reluctant = !reluctant
	}
	exp := runes[0:l]
	switch runes[0] {
	case '?':
		return repeatNode{
			Min:       0,
			Max:       1,
			Reluctant: reluctant,
			Atomic:    atomic,
			Exp:       exp,
		}, l
	case '*':
		return repeatNode{
			Min:       0,
			Max:       -1,
			Reluctant: reluctant,
			Atomic:    atomic,
			Exp:       exp,
		}, l
	case '+':
		return repeatNode{
			Min:       1,
			Max:       -1,
			Reluctant: reluctant,
			Atomic:    atomic,
			Exp:       exp,
		}, l
	}
	panic("unreachable")
	return nil, 0
}

func (p *parser) fetchGroup(runes []rune, flags syntax.Flags) (node, int) {
	g := 1
	meta := false
	for i, r := range runes[1:] {
		if meta {
			meta = false
			continue
		}
		switch r {
		case '\\':
			meta = true
		case '(':
			if meta {
				meta = false
			} else {
				g++
			}
		case ')':
			if meta {
				meta = false
			} else {
				g--
				if g == 0 {
					return p.group(runes[1:i+1], flags), i + 2
				}
			}
		default:
			if meta {
				meta = false
			}
		}
	}
	return nil, 0
}

func (p *parser) concatRepetitions(nodes []node) []node {
	r := []node{}
	for _, n := range nodes {
		if rn, ok := n.(repeatNode); ok {
			if len(r) == 0 {
				panic(newErrorRunes(syntax.ErrMissingRepeatArgument, rn.Exp))
			}
			if nrn, ok := r[len(r)-1].(repeatNode); ok {
				panic(newErrorRunes(syntax.ErrInvalidRepeatOp, append(nrn.Exp, rn.Exp...)))
			}
			rn.N = r[len(r)-1]
			if any, ok := rn.N.(anyCharNode); ok {
				r[len(r)-1] = anyCharRepeatNode{
					Flags:     any.Flags,
					Min:       rn.Min,
					Max:       rn.Max,
					Reluctant: rn.Reluctant,
					Atomic:    rn.Atomic,
				}
			} else {
				r[len(r)-1] = rn
			}
			continue
		}
		r = append(r, n)
	}
	return r
}

func (p *parser) concatLiterals(nodes []node) []node {
	r := []node{}
	for _, n := range nodes {
		if sln, ok := n.(literalNode); ok {
			if len(r) > 0 {
				if fln, ok := r[len(r)-1].(literalNode); ok {
					if sln.Flags == fln.Flags {
						fln.L = append(fln.L, sln.L...)
						r[len(r)-1] = fln
						continue
					}
				}
			}
		}
		r = append(r, n)
	}
	return r
}

func (p *parser) concatAlternations(nodes []node) []node {
	altidx := []int{0}
	res := []node{}
	for i, n := range nodes {
		if _, ok := n.(alterNode); ok {
			altidx = append(altidx, i, i+1)
		}
	}
	altidx = append(altidx, len(nodes))
	if len(altidx) <= 2 {
		return nodes
	}
	for i := 0; i < len(altidx)/2; i++ {
		b := altidx[i*2]
		e := altidx[i*2+1]
		g := nodes[b:e]
		if len(g) == 0 {
			res = append(res, nil)
		} else if len(g) == 1 {
			res = append(res, g[0])
		} else {
			res = append(res, groupNode{N: g})
		}
	}
	uniq := make([]node, 0, len(res))
loop:
	for _, n := range res {
		for _, d := range uniq {
			if reflect.DeepEqual(n, d) {
				continue loop
			}
		}
		uniq = append(uniq, n)
	}
	return []node{alterNode{N: uniq}}
}

func (p *parser) removeSequentialBoundaries(nodes []node) []node {
	r := []node{}
	for _, n := range nodes {
		if len(r) > 0 {
			switch n.(type) {
			case beginNode, endNode, wordBoundaryNode:
				if reflect.DeepEqual(n, r[len(r)-1]) {
					continue
				}
			}
		}
		r = append(r, n)
	}
	return r
}
