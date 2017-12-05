package parser

import (
    "strconv"
    "log"
)

type Parser struct {
    BaseStr string
    curPos int
}

func (p *Parser) Exec() int {
    return p.number()
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
    for {
        char := p.peek()
        _, err := strconv.Atoi(string(char))
        if err != nil {
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



