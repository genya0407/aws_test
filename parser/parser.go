package parser

import (
    "strconv"
    "log"
)

func Parse(str string) int {
    p := Parser { BaseStr: str }
    return p.Exec()
}

type Parser struct {
    BaseStr string
    curPos int
}

func (p *Parser) Exec() int {
    return p.expr()
}

func (p *Parser) peek() byte {
    if len(p.BaseStr) > p.curPos {
        return p.BaseStr[p.curPos]
    } else {
        return 0
    }
}

func (p *Parser) next() {
    p.curPos += 1
}

func (p *Parser) number() int {
    bs := []byte{}
    if p.peek() == '-' {
        bs = append(bs, '-')
        p.next()
    }

    for {
        char := p.peek()
        if !isDigit(char) {
            break
        }
        bs = append(bs, char)
        p.next()
    }
    num, err := strconv.Atoi(string(bs))
    if err != nil {
        log.Fatal("Parser#number errors")
    }
    return num
}

// expr = term, [{+|-} term]
func (p *Parser) expr() int {
    x := p.term()
    for {
        switch(p.peek()) {
        case '+':
            p.next()
            x += p.term()
            continue
        case '-':
            p.next()
            x -= p.term()
            continue
        }
        break
    }
    return x
}

// term = factor, [{*|/}, factor]
func (p *Parser) term() int {
    x := p.factor()

    for {
        switch (p.peek()) {
        case '*':
            p.next()
            x *= p.factor()
            continue
        case '/':
            p.next()
            x /= p.factor()
            continue;
        }
        break
    }
    return x;
}

// factor = (, expr, ) | (, number, ) | number
func (p *Parser) factor() int {
    if (p.peek() == '(') {
        p.next()
        x := p.expr()
        if (p.peek() == ')') {
            p.next()
        }
        return x
    } else {
        return p.number()
    }
}

func isDigit(char byte) bool {
    _, err := strconv.Atoi(string(char))
    if err != nil {
        return false
    } else {
        return true
    }
}

