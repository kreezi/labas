package tokenizer

import (
	. "ilya_golang/Laba32/types"
	"unicode"
)

type Token struct {
	Type, Value string
}

func Tokenize(expression string) []Token {
	var tokens []Token

	expressionLen := len(expression)
	for i := 0; i < expressionLen; {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) {
			start := i
			for i < expressionLen && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			tokens = append(tokens, Token{Type: NUMBER, Value: expression[start:i]})
		} else if unicode.IsLetter(ch) {
			start := i
			for i < expressionLen && (unicode.IsLetter(rune(expression[i])) || unicode.IsDigit(rune(expression[i]))) {
				i++
			}
			value := expression[start:i]

			if i < expressionLen && expression[i] == '(' {
				tokens = append(tokens, Token{Type: FUNCTION, Value: value})
			} else {
				tokens = append(tokens, Token{Type: IDENT, Value: expression[start:i]})
			}
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == ',' {
			tokens = append(tokens, Token{Type: OPERATOR, Value: string(ch)})
			i++
		} else if ch == '(' {
			tokens = append(tokens, Token{Type: LPAREN, Value: string(ch)})
			i++
		} else if ch == ')' {
			tokens = append(tokens, Token{Type: RPAREN, Value: string(ch)})
			i++
		} else {
			// Пропускаем пробелы и другие незначащие символы
			i++
		}
	}
	return tokens
}
