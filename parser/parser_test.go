package parser

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/genya0407/aws/parser"
)

func TestParser(t *testing.T) {
    check(t, "11", 11)
    check(t, "(11+100)", 111)
    check(t, "(11-100)", -89)
    check(t, "-100", -100)
    check(t, "3*2", 6)
    check(t, "4/2", 2)
    check(t, "1+2*3", 7)
    check(t, "(1+2)*3", 9)
}

func check(t *testing.T, baseStr string, expected int) {
    p := parser.Parser { BaseStr: baseStr }
    assert.Equal(t, expected, p.Exec())
}