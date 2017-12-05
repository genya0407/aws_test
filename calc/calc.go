package calc

import (
    "strings"
)

func Calc(query string) string {
    if isValidQuery(query) {
        return fmt.Sprintf("%d", calc(query))
    } else {
        return "ERROR"
    }
}

const validChars = "*/+-()0123456789"
func isValidQuery(query string) bool {
    // start with number
    if !strings.Contains(numbers, query[0]) {
        return false
    }

    // only number or operators
    for _, char := range query {
        if !strings.Contains(validChars, char) {
            return false
        }
    }

    return true
}

/* Core logic */



func calc(query string) int {

}