package main

import (
	"fmt"
	"strings"
)

func removeSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

func match(opening, closing rune) bool {
	return (opening == '(' && closing == ')') || (opening == '[' && closing == ']') || (opening == '{' && closing == '}')
}

func main() {
	var s string
	fmt.Scan(&s)

	s = removeSpaces(s)

	stack := []rune{}
	isBalanced := true

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 {
				isBalanced = false
				break
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if !match(top, char) {
				isBalanced = false
				break
			}
		}
	}

	if len(stack) > 0 {
		isBalanced = false
	}

	if isBalanced {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	fmt.Println()

}
