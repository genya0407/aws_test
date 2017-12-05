package parser

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/genya0407/aws/parser"
)

func TestParser(t *testing.T) {
    p := parser.Parser { BaseStr: "11" }
    assert.Equal(t, 11, p.Exec())
}
