package metatools

import "strings"

type indextion struct {
	I int
	W string
}

// Parser fix major issues of olivia
type Parser struct {
	data string
	sd   []indextion
}

// Init the parser
func (p *Parser) Init(d string) {
	p.data = strings.ToLower(d)
	var index indextion

	var whitespace bool = true
	var wordstart int = 0
	for i, c := range p.data {
		if c != ' ' {
			whitespace = false
			continue
		}
		if whitespace == true {
			wordstart = i + 1
			continue
		}

		index.I = wordstart
		index.W = p.data[wordstart:i]
		p.sd = append(p.sd, index)
		whitespace = true
		wordstart = i + 1
	}
	index.I = wordstart
	index.W = p.data[wordstart:]
	p.sd = append(p.sd, index)
}

func (p *Parser) count(needle string) int {
	count := 0
	for _, it := range p.sd {
		if it.W == needle {
			count++
		}
	}
	return count
}

func (p *Parser) beforeIdx(needle string, x int) int {
	if x < 0 {
		x = p.count(needle) + x
	}

	countdown := 0

	for i, it := range p.sd {
		if needle == it.W {
			if countdown == x {
				if p.sd[i].I == 0 {
					return 0
				}
				return p.sd[i].I - 1
			}
			countdown++
		}
	}
	return 0
}

func (p *Parser) afterIdx(needle string, x int) int {
	if x < 0 {
		x = p.count(needle) + x
	}

	countdown := 0

	for i, it := range p.sd {
		if needle == it.W {
			if countdown == x {
				if i == len(p.sd)-1 {
					return len(p.data)
				}
				return p.sd[i+1].I
			}
			countdown++
		}
	}
	return len(p.data)
}

// After return the rest of the string after the last occurence of needle
func (p *Parser) After(needle string, x int) string {
	return p.data[p.afterIdx(needle, x):]
}

// Before return the rest of the string before the last occurence of needle
func (p *Parser) Before(needle string, x int) string {
	return p.data[:p.beforeIdx(needle, x)]
}

// Between returns the first occurence of a substring in middle of start - end
func (p *Parser) Between(start, end string, x, y int) string {
	aft := p.afterIdx(start, x)
	bfr := p.beforeIdx(end, y)
	if aft > bfr {
		return ""
	}
	return p.data[aft:bfr]
}
