package calc

import (
    . "github.com/genya0407/aws/calc"
    "testing"
    "github.com/stretchr/testify/assert"
    "time"
)

func TestCalc(t *testing.T) {
    exec(t, "abc", "ERROR")
    exec(t, "1+1", "2")
    exec(t, "2-1", "1")
    exec(t, "3*2", "6")
    exec(t, "4/2", "2")
    exec(t, "1+2*3", "7")
    exec(t, "(1+2)*3", "9")
}

func exec(t *testing.T, query string, expected string) {
    assert.Equal(t, expected, calc.Calc(query))
}