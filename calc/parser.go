package calc

/*
    Query should be parsed as Expr.
    Expr = Formula | Value
    Formula = Expr Operator Expr
    Value = int
*/


type Operator int
const (
    PLUS     Operator = iota
    MINUS    Operator = iota
    MULTIPLY Operator = iota
    DIVIDE   Operator = iota
)

type Expr interface {
    Result() int
}

type Value int
func (self *Value) Result() int {
    return self
}

type Formula struct {
    LeftOperand Operand
    RightOperand Operand
    Operator Operator
}
func (self *Formula) Result() int {
    left := self.LeftOperand.Val()
    right := self.RightOperand.Val()

    switch self.Operator {
    case PLUS:
        return left + right
    case MINUS:
        return left - right
    case MULTIPLY:
        return left * right
    case DIVIDE:
        return left / right
    }
}