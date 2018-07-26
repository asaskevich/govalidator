package syntax

import (
	"bytes"
	"errors"
	"regexp/syntax"
	"unicode/utf8"
)

var errDeadFiber = errors.New("Resume dead fiber")

const (
	hintFixedBeginning = iota
)

type input struct {
	b, o  []byte
	begin int
	sub   submatch
	funcs []FuncMap
}

func (i input) Substr(offset int, sub submatch) input {
	if offset > len(i.b) {
		offset = len(i.b)
	}
	return input{
		b:     i.b[offset:],
		o:     i.o,
		begin: i.begin + offset,
		sub:   sub,
		funcs: i.funcs,
	}
}

type output struct {
	offset int
	sub    submatch
}

type matchLocation struct {
	begin int
	b     []byte
}

type submatch struct {
	i map[int]matchLocation
	n map[string]matchLocation
}

func (s submatch) Merge(m submatch) submatch {
	i := make(map[int]matchLocation, len(s.i)+len(m.i))
	n := make(map[string]matchLocation, len(s.n)+len(m.n))
	for k, v := range s.i {
		i[k] = v
	}
	for k, v := range s.n {
		n[k] = v
	}
	for k, v := range m.i {
		i[k] = v
	}
	for k, v := range m.n {
		n[k] = v
	}
	return submatch{
		i: i,
		n: n,
	}
}

type fiberOutput struct {
	output
	err error
}

type fiber interface {
	Resume() (output, error)
}

type hint map[int]interface{}

type node interface {
	Fiber(i input) fiber
	IsExtended() bool
	MinMax() (int, int)
	LiteralPrefix() ([]byte, bool)
	Hint() hint
}

type minmax [2]int

// flagNode represents a flag expression: /(?i)/
type flagNode struct {
	Flags map[syntax.Flags]int
}

func (n flagNode) Fiber(i input) fiber {
	panic("pseudo node")
	return nil
}

func (n flagNode) IsExtended() bool {
	return false
}

func (n flagNode) MinMax() (int, int) {
	return 0, 0
}

func (n flagNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n flagNode) Hint() hint {
	return nil
}

// groupNode represents a group expression: /([exp])/
type groupNode struct {
	N          []node
	Atomic     bool
	Index      int
	Name       string
	Repetition int
	mm         *minmax
}

func (n groupNode) size() int {
	size := len(n.N)
	if size == 1 && n.Repetition > 0 {
		size = n.Repetition
	}
	return size
}

func (n groupNode) Fiber(i input) fiber {
	return &groupNodeFiber{
		I:      i,
		node:   n,
		stack:  make([]*output, n.size()),
		fstack: make([]fiber, n.size()),
	}
}

func (n groupNode) IsExtended() bool {
	if n.Atomic {
		return true
	}
	for _, e := range n.N {
		if e.IsExtended() {
			return true
		}
	}
	return false
}

func (n groupNode) IsAnonymous() bool {
	return n.Index == 0 && len(n.Name) == 0
}

func (n groupNode) LiteralPrefix() ([]byte, bool) {
	if len(n.N) == 0 {
		return nil, true
	}
	var b []byte
	comp := true
	for _, e := range n.N {
		p, c := e.LiteralPrefix()
		b = append(b, p...)
		if !c {
			comp = false
			break
		}
	}
	return b, comp
}

func (n groupNode) MinMax() (int, int) {
	if n.mm != nil {
		return n.mm[0], n.mm[1]
	}
	gmin := 0
	gmax := 0

	s := n.size()
	for i := 0; i < s; i++ {
		e := n.N[0]
		if n.Repetition == 0 {
			e = n.N[i]
		}
		min, max := e.MinMax()
		if gmin >= 0 {
			gmin += min
		}
		if max < 0 {
			gmax = -1
		} else if gmax >= 0 {
			gmax += max
		}
	}
	n.mm = &minmax{gmin, gmax}
	return gmin, gmax
}

func (n groupNode) Hint() hint {
	for _, e := range n.N {
		if b, _ := e.Hint()[hintFixedBeginning].(bool); b {
			return hint{hintFixedBeginning: true}
		}
	}
	return nil
}

type groupNodeFiber struct {
	I      input
	node   groupNode
	stack  []*output
	fstack []fiber
	fixed  bool
}

func (f *groupNodeFiber) Resume() (output, error) {
	if f.fixed {
		return output{}, errDeadFiber
	}

	min, _ := f.node.MinMax()
	if len(f.I.b) < min {
		f.fixed = true
		return output{}, errDeadFiber
	}

	if len(f.node.N) == 0 {
		f.fixed = true
		return output{
			offset: 0,
		}, nil
	}

mainloop:
	for {
		offset := 0
		var s submatch
		s = s.Merge(f.I.sub)
		size := f.node.size()
	stloop:
		for i := 0; i < size; i++ {
			n := f.node.N[0]
			if f.node.Repetition == 0 {
				n = f.node.N[i]
			}
			if f.fstack[i] == nil {
				f.fstack[i] = n.Fiber(f.I.Substr(offset, s))
			}
			if f.stack[i] == nil {
				o, err := f.fstack[i].Resume()
				if err != nil {
					if i == 0 {
						// no match
						break mainloop
					} else {
						// backtrack
						f.fstack[i] = nil
						f.stack[i-1] = nil
					}
					break stloop
				} else {
					if i >= size-1 {
						b := f.I.b[:offset+o.offset]

						s = s.Merge(submatch{})
						if f.node.Index > 0 {
							s.i[f.node.Index] = matchLocation{begin: f.I.begin, b: b}
						}
						if len(f.node.Name) > 0 {
							s.n[f.node.Name] = matchLocation{begin: f.I.begin, b: b}
						}

						if f.node.Atomic {
							f.fixed = true
						}
						return output{
							offset: len(b),
							sub:    s.Merge(o.sub),
						}, nil
					}
					f.stack[i] = &o
				}
			}
			offset += f.stack[i].offset
			s = s.Merge(f.stack[i].sub)
		}
	}

	return output{}, errDeadFiber
}

type anyCharRepeatNode struct {
	Flags     syntax.Flags
	Min, Max  int
	Reluctant bool
	Atomic    bool
}

func (n anyCharRepeatNode) IsExtended() bool {
	return n.Atomic
}

func (n anyCharRepeatNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n anyCharRepeatNode) MinMax() (int, int) {
	max := -1
	if n.Max >= 0 {
		max = utf8.UTFMax * n.Max
	}
	return 1 * n.Min, max
}

func (n anyCharRepeatNode) Fiber(i input) fiber {
	return &anyCharRepeatNodeFiber{I: i, node: n}
}

func (n anyCharRepeatNode) Hint() hint {
	return nil
}

type anyCharRepeatNodeFiber struct {
	I      input
	node   anyCharRepeatNode
	b      []byte
	cnt    int
	runes  int
	offset int
	fixed  bool
}

func (f *anyCharRepeatNodeFiber) Resume() (output, error) {
	if f.fixed {
		return output{}, errDeadFiber
	}
	if f.cnt == 0 {
		f.b = f.I.b
		if f.node.Flags&syntax.DotNL == 0 {
			i := bytes.IndexByte(f.b, '\n')
			if i >= 0 {
				f.b = f.b[:i]
			}
		}
		if f.node.Reluctant {
			for i := 0; i < f.node.Min; i++ {
				_, n := utf8.DecodeRune(f.b[f.offset:])
				if n == 0 {
					break
				}
				f.offset += n
				f.runes++
			}
		} else {
			if f.node.Max < 0 {
				f.runes = len(bytes.Runes(f.b))
				f.offset = len(f.b)
			} else {
				for len(f.b[f.offset:]) > 0 && f.runes < f.node.Max {
					_, n := utf8.DecodeRune(f.b[f.offset:])
					if n == 0 {
						break
					}
					f.offset += n
					f.runes++
				}
			}
		}
		f.cnt++
	}

	if f.runes < f.node.Min || (f.node.Max >= 0 && f.runes > f.node.Max) {
		return output{}, errDeadFiber
	}

	if f.node.Atomic {
		f.fixed = true
	}
	o := output{offset: f.offset, sub: f.I.sub}
	if f.node.Reluctant {
		_, n := utf8.DecodeRune(f.b[f.offset:])
		if n == 0 {
			f.fixed = true
		}
		f.offset += n
		f.runes++
	} else {
		_, n := utf8.DecodeLastRune(f.b[:f.offset])
		if n == 0 {
			f.fixed = true
		}
		f.offset -= n
		f.runes--
	}
	return o, nil
}

// repeatNode represents a repeat expression: /[exp]+/
type repeatNode struct {
	N         node
	Min, Max  int
	Reluctant bool
	Atomic    bool
	Exp       []rune
}

func (n repeatNode) Fiber(i input) fiber {

	f := repeatNodeFiber{I: i, node: n}

	max := f.node.Max
	min, _ := f.node.MinMax()

	if max < 0 {
		if min <= 0 {
			max = len(f.I.b)
		} else {
			max = len(f.I.b) / min
		}
	}

	for i := min; i <= max; i++ {
		g := groupNode{N: []node{n.N}, Repetition: i}
		if i == 0 {
			g.N = []node(nil)
		}
		gf := g.Fiber(f.I.Substr(0, f.I.sub))
		_, err := gf.Resume()
		if err != nil {
			max = i - 1
			break
		}
	}

	if max < f.node.Min {
		f.s = 0
		f.e = 0
		f.suc = 0
	} else if f.node.Reluctant {
		f.s = f.node.Min
		f.e = max + 1
		f.suc = 1
	} else {
		f.s = max
		f.e = f.node.Min - 1
		f.suc = -1
	}
	f.cnt = f.s
	return &f
}

func (n repeatNode) IsExtended() bool {
	return n.Atomic || n.N.IsExtended()
}

func (n repeatNode) LiteralPrefix() ([]byte, bool) {
	if n.Min == 0 {
		if n.Max == 0 {
			return nil, true
		}
		return nil, false
	}
	p, comp := n.N.LiteralPrefix()
	if !comp {
		return p, false
	}
	return bytes.Repeat(p, n.Min), (n.Min == n.Max)
}

func (n repeatNode) MinMax() (int, int) {
	min, max := n.N.MinMax()
	rmin := n.Min * min
	rmax := 0
	if max < 0 || n.Max < 0 {
		rmax = -1
	} else {
		rmax = max * n.Max
	}
	return rmin, rmax
}

func (n repeatNode) Hint() hint {
	if b, _ := n.N.Hint()[hintFixedBeginning].(bool); !b {
		return nil
	}
	if n.Min > 0 {
		return hint{hintFixedBeginning: true}
	}
	return nil
}

type repeatNodeFiber struct {
	I         input
	node      repeatNode
	s, e, suc int
	cnt       int
	group     fiber
	fixed     bool
}

func (f *repeatNodeFiber) Resume() (output, error) {
	if f.fixed {
		return output{}, errDeadFiber
	}

loop:
	for f.cnt != f.e {
		if f.cnt == 0 {
			f.group = nil
			f.cnt += f.suc
			if f.node.Atomic {
				f.fixed = true
			}
			return output{
				offset: 0,
				sub:    f.I.sub,
			}, nil
		}

		if f.group == nil {
			n := groupNode{N: []node{f.node.N}, Repetition: f.cnt}
			if f.cnt == 0 {
				n.N = []node(nil)
			}
			f.group = n.Fiber(f.I.Substr(0, f.I.sub))
		}

		o, err := f.group.Resume()
		if err != nil {
			f.group = nil
			f.cnt += f.suc
			continue loop
		}

		if f.node.Atomic {
			f.fixed = true
		}
		return output{
			offset: o.offset,
			sub:    o.sub,
		}, nil
	}
	return output{}, errDeadFiber
}

// alterNode represents an alternation expression: /[exp]|[exp]/
type alterNode struct {
	N []node
}

func (n alterNode) Fiber(i input) fiber {
	fibers := make([]fiber, len(n.N))
	for index, n := range n.N {
		if n != nil {
			fibers[index] = n.Fiber(i)
		}
	}
	return &alterNodeFiber{I: i, node: n, fibers: fibers}
}

func (n alterNode) IsExtended() bool {
	for _, e := range n.N {
		if e != nil && e.IsExtended() {
			return true
		}
	}
	return false
}

func (n alterNode) LiteralPrefix() ([]byte, bool) {
	if len(n.N) == 0 {
		return nil, false
	}
	var b [][]byte
	minlen := -1
	for _, e := range n.N {
		if e == nil {
			minlen = 0
		} else {
			p, _ := e.LiteralPrefix()
			b = append(b, p)
			if minlen < 0 || minlen > len(p) {
				minlen = len(p)
			}
		}
	}
	if minlen < 0 {
		minlen = 0
	}
	i := 0
loop:
	for ; i < minlen; i++ {
		for _, p := range b {
			if !bytes.HasPrefix(p, b[0][:i+1]) {
				break loop
			}
		}
	}
	return b[0][:i], false
}

func (n alterNode) MinMax() (int, int) {
	amin := -1
	amax := 0
	for _, e := range n.N {
		min := 0
		max := 0
		if e != nil {
			min, max = e.MinMax()
		}
		if amin < 0 || amin > min {
			amin = min
		}
		if max < 0 {
			amax = -1
		} else if amax >= 0 && amax < max {
			amax = max
		}
	}
	return amin, amax
}

func (n alterNode) Hint() hint {
	for _, e := range n.N {
		if e == nil {
			return nil
		}
		if b, _ := e.Hint()[hintFixedBeginning].(bool); !b {
			return nil
		}
	}
	if len(n.N) > 0 {
		return hint{hintFixedBeginning: true}
	}
	return nil
}

type alterNodeFiber struct {
	I      input
	node   alterNode
	fibers []fiber
	cnt    int
}

func (f *alterNodeFiber) Resume() (output, error) {
	for f.cnt < len(f.node.N) {
		if f.fibers[f.cnt] == nil {
			f.cnt++
			return output{offset: 0}, nil
		} else if o, err := f.fibers[f.cnt].Resume(); err == nil {
			return output{offset: o.offset, sub: o.sub}, nil
		} else {
			f.cnt++
		}
	}
	return output{}, errDeadFiber
}

type anyCharNode struct {
	Flags    syntax.Flags
	Reversed bool
}

func (n anyCharNode) Fiber(i input) fiber {
	return &anyCharNodeFiber{I: i, node: n}
}

func (n anyCharNode) IsExtended() bool {
	return false
}

func (n anyCharNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n anyCharNode) MinMax() (int, int) {
	return 1, utf8.UTFMax
}

func (n anyCharNode) Hint() hint {
	return nil
}

type anyCharNodeFiber struct {
	I    input
	node anyCharNode
	cnt  int
}

func (f *anyCharNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		r, size := utf8.DecodeRune(f.I.b)
		if f.node.Flags&syntax.DotNL == 0 && r == '\n' {
			return output{}, errDeadFiber
		}
		if size > 0 {
			return output{offset: size}, nil
		}
	}
	return output{}, errDeadFiber
}

// charNode represents a character expression: /[a-z]/
type charNode struct {
	Flags    syntax.Flags
	Matcher  []charNodeMatcher
	Reversed bool
}

func (n charNode) Fiber(i input) fiber {
	return &charNodeFiber{I: i, node: n}
}

func (n charNode) IsExtended() bool {
	return false
}

func (n charNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n charNode) MinMax() (int, int) {
	return 1, utf8.UTFMax
}

func (n charNode) Hint() hint {
	return nil
}

type charNodeFiber struct {
	I    input
	node charNode
	cnt  int
}

func (f *charNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		r, size := utf8.DecodeRune(f.I.b)
		if size > 0 {
			m := false
			for _, mf := range f.node.Matcher {
				if mf.Match(r, f.node.Flags) {
					m = true
					break
				}
			}
			if f.node.Reversed {
				m = !m
			}
			if m {
				return output{offset: size}, nil
			}
		}
	}
	return output{}, errDeadFiber
}

// literalNode represents a literal expression: /string/
type literalNode struct {
	Flags syntax.Flags
	L     []byte
}

func (n literalNode) Fiber(i input) fiber {
	return &literalNodeFiber{I: i, node: n}
}

func (n literalNode) IsExtended() bool {
	return false
}

func (n literalNode) LiteralPrefix() ([]byte, bool) {
	if n.Flags&syntax.FoldCase != 0 {
		return nil, false
	}
	return n.L, true
}

func (n literalNode) MinMax() (int, int) {
	return len(n.L), len(n.L)
}

func (n literalNode) Hint() hint {
	return nil
}

type literalNodeFiber struct {
	I    input
	node literalNode
	cnt  int
}

func (f *literalNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++

		l := len(f.node.L)
		if l > len(f.I.b) {
			l = len(f.I.b)
		}
		if f.node.Flags&syntax.FoldCase != 0 && bytes.EqualFold(f.node.L, f.I.b[:l]) {
			return output{offset: l}, nil
		} else if bytes.Equal(f.node.L, f.I.b[:l]) {
			return output{offset: l}, nil
		}
	}
	return output{}, errDeadFiber
}

// beginNode represents a begginning expression: /^/
type beginNode struct {
	Flags syntax.Flags
	Line  bool
}

func (n beginNode) Fiber(i input) fiber {
	return &beginNodeFiber{I: i, node: n}
}

func (n beginNode) IsExtended() bool {
	return false
}

func (n beginNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n beginNode) MinMax() (int, int) {
	return 0, 0
}

func (n beginNode) Hint() hint {
	if n.Flags&syntax.OneLine == 0 {
		return nil
	}
	return hint{hintFixedBeginning: true}
}

type beginNodeFiber struct {
	I    input
	node beginNode
	cnt  int
}

func (f *beginNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		if f.I.begin == 0 {
			return output{offset: 0}, nil
		}
		if f.node.Line && f.node.Flags&syntax.OneLine == 0 && f.I.o[f.I.begin-1] == '\n' {
			return output{offset: 0}, nil
		}
	}
	return output{}, errDeadFiber
}

// endNode represents an end expression: /$/
type endNode struct {
	Flags syntax.Flags
	Line  bool
}

func (n endNode) Fiber(i input) fiber {
	return &endNodeFiber{I: i, node: n}
}

func (n endNode) IsExtended() bool {
	return false
}

func (n endNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n endNode) MinMax() (int, int) {
	return 0, 0
}

func (n endNode) Hint() hint {
	return nil
}

type endNodeFiber struct {
	I    input
	node endNode
	cnt  int
}

func (f *endNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		if len(f.I.b) == 0 {
			return output{offset: 0}, nil
		}
		if f.node.Line && f.node.Flags&syntax.OneLine == 0 && len(f.I.b) > 0 && f.I.b[0] == '\n' {
			return output{offset: 0}, nil
		}
	}
	return output{}, errDeadFiber
}

type wordBoundaryNode struct {
	Reversed bool
}

func (n wordBoundaryNode) Fiber(i input) fiber {
	return &wordBoundaryFiber{I: i, node: n}
}

func (n wordBoundaryNode) IsExtended() bool {
	return false
}

func (n wordBoundaryNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n wordBoundaryNode) MinMax() (int, int) {
	return 0, 0
}

func (n wordBoundaryNode) Hint() hint {
	return nil
}

type wordBoundaryFiber struct {
	I    input
	node wordBoundaryNode
	cnt  int
}

func (f *wordBoundaryFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		match := false
		if len(f.I.b) > 0 {
			if f.I.begin > 0 && isASCIIWord(rune(f.I.b[0])) != isASCIIWord(rune(f.I.o[f.I.begin-1])) {
				match = true
			}
			if f.I.begin == 0 && isASCIIWord(rune(f.I.b[0])) {
				match = true
			}
		}
		if len(f.I.o) > 0 && f.I.begin == len(f.I.o) {
			r, _ := utf8.DecodeLastRune(f.I.o)
			if isASCIIWord(r) {
				match = true
			}
		}
		if f.node.Reversed {
			match = !match
		}
		if match {
			return output{offset: 0}, nil
		}
	}
	return output{}, errDeadFiber
}

// backRefNode represents a back reference expression: /\1/
type backRefNode struct {
	Flags syntax.Flags
	Index int
	Name  string
}

func (n backRefNode) Fiber(i input) fiber {
	return &backRefNodeFiber{I: i, node: n}
}

func (n backRefNode) IsExtended() bool {
	return true
}

func (n backRefNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n backRefNode) MinMax() (int, int) {
	return 0, -1
}

func (n backRefNode) Hint() hint {
	return nil
}

type backRefNodeFiber struct {
	I    input
	node backRefNode
	cnt  int
}

func (f *backRefNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++

		var b []byte
		if f.node.Index > 0 {
			if r, ok := f.I.sub.i[f.node.Index]; ok {
				b = r.b
			}
		} else if len(f.node.Name) > 0 {
			if r, ok := f.I.sub.n[f.node.Name]; ok {
				b = r.b
			}
		}

		l := len(b)
		if l > len(f.I.b) {
			l = len(f.I.b)
		}
		if f.node.Flags&syntax.FoldCase != 0 && bytes.EqualFold(b, f.I.b[:l]) {
			return output{offset: l}, nil
		} else if bytes.Equal(b, f.I.b[:l]) {
			return output{offset: l}, nil
		}
	}
	return output{}, errDeadFiber
}

type lookaheadNode struct {
	N        node
	Negative bool
}

func (n lookaheadNode) Fiber(i input) fiber {
	return &lookaheadNodeFiber{I: i, node: n}
}

func (n lookaheadNode) IsExtended() bool {
	return true
}

func (n lookaheadNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n lookaheadNode) MinMax() (int, int) {
	return 0, -1
}

func (n lookaheadNode) Hint() hint {
	return nil
}

type lookaheadNodeFiber struct {
	I    input
	node lookaheadNode
	cnt  int
}

func (f *lookaheadNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		_, err := f.node.N.Fiber(f.I).Resume()
		if (err == nil) != f.node.Negative {
			return output{offset: 0}, nil
		}
	}
	return output{}, errDeadFiber
}

type lookbehindNode struct {
	N        node
	Negative bool
}

func (n lookbehindNode) Fiber(i input) fiber {
	return &lookbehindNodeFiber{I: i, node: n}
}

func (n lookbehindNode) IsExtended() bool {
	return true
}

func (n lookbehindNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n lookbehindNode) MinMax() (int, int) {
	return 0, -1
}

func (n lookbehindNode) Hint() hint {
	return nil
}

type lookbehindNodeFiber struct {
	I    input
	node lookbehindNode
	cnt  int
}

func (f *lookbehindNodeFiber) Resume() (output, error) {
	if f.cnt == 0 {
		f.cnt++
		min, max := f.node.N.MinMax()
		if max < 0 {
			panic("Lookbehind only supports a finite length matching")
		}
		if max > f.I.begin {
			max = f.I.begin
		}
		if min <= max {
			match := false
			for i := min; i <= max; i++ {
				in := input{
					b:     f.I.o[f.I.begin-i:],
					o:     f.I.o,
					begin: f.I.begin - i,
					sub:   f.I.sub,
				}
				_, err := f.node.N.Fiber(in).Resume()
				if err == nil {
					match = true
					break
				}
			}
			if f.node.Negative {
				match = !match
			}
			if match {
				return output{offset: 0}, nil
			}
		}
	}
	return output{}, errDeadFiber
}

type funcNode struct {
	Name string
}

func (n funcNode) Fiber(i input) fiber {
	return &funcNodeFiber{I: i, node: n}
}

func (n funcNode) IsExtended() bool {
	return true
}

func (n funcNode) LiteralPrefix() ([]byte, bool) {
	return nil, false
}

func (n funcNode) MinMax() (int, int) {
	return 0, -1
}

func (n funcNode) Hint() hint {
	return nil
}

type funcNodeFiber struct {
	I    input
	node funcNode
}

func (f *funcNodeFiber) Resume() (output, error) {
	matches := make(map[interface{}][]int)
	for k, v := range f.I.sub.i {
		matches[k] = []int{v.begin, v.begin + len(v.b)}
	}
	for k, v := range f.I.sub.n {
		matches[k] = []int{v.begin, v.begin + len(v.b)}
	}
	for _, m := range f.I.funcs {
		if fun, ok := m[f.node.Name]; ok {
			res := fun(Context{
				Data:    f.I.o,
				Cursor:  f.I.begin,
				Matches: matches,
			})
			if res == nil {
				return output{offset: 0}, nil
			}
			switch v := res.(type) {
			case nil:
				return output{offset: 0}, nil
			case int:
				if v >= 0 {
					return output{offset: v}, nil
				}
				return output{}, errDeadFiber
			}
		}
	}
	return output{}, errDeadFiber
}
