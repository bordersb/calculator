// To execute Go code, please declare a func main() in a package "main" // Write a function, calculate(), that accepts a string of addition / subtraction operations as an argument. The function should return the result of the operations as an integer.

// ex: calculate("1 - 2 + 3") // 2

// Next, revise your calculate function so that it accepts a string of addition / subtraction operations and also parentheses to indicate order of operations.

// ex: calculate("1 - (2 + 3)") // -4

package main

import (
	"fmt"
)

// findClosingParenIdx finds the closing parenthesis associated with the
// starting parenthesis in the string
func findClosingParenIdx(s string, start int) int {
	depth := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {

				return i
			}
		}
	}

	return len(s) - 1
}

func calculate(s string) (int, error) {
	sum := 0
	sign := 1
	curr := 0
	prev := ""

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == ' ':
			continue
		case s[i] == '(':
			end := findClosingParenIdx(s, i)
			var err error
			curr, err = calculate(s[i+1 : end])
			if err != nil {
				return -1, err
			}
			i = end
		case isDigit(s[i]):
			curr = curr*10 + (int(s[i]) - '0')
		case s[i] == '+', s[i] == '-':
			sum += sign * curr
			curr = 0
			if s[i] == '+' {
				sign = 1
			}
			if s[i] == '-' {
				if prev == "-" {
					sign = 1
				} else {
					sign = -1
				}
			}
		default:
			return -1, fmt.Errorf("unknown char %s", s[i])
		}
		prev = string(s[i])
	}

	// finish calculating what is left over in curr
	sum += sign * curr
	return sum, nil
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func main() {
	n, _ := calculate("1 - (2 + 3)")
	fmt.Printf("number returned  %d \n", n)

	n2, _ := calculate("1 + 2 -3")
	fmt.Printf("number returned  %d \n", n2)

	n3, _ := calculate("1 - (2 - 30)")
	fmt.Printf("number returned  %d \n", n3)
}
