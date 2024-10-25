package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenType string

const (
	TokenIdentifier TokenType = "IDENT"
	TokenNumber     TokenType = "NUMBER"
	TokenOperator   TokenType = "OP"
	TokenAssign     TokenType = "ASSIGN"
	TokenEOF        TokenType = "EOF"
	TokenError      TokenType = "ERROR"
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) []Token {
	var tokens []Token

	// Define regular expressions for different tokens
	tokenSpec := []struct {
		typ  TokenType
		regx *regexp.Regexp
	}{
		{TokenNumber, regexp.MustCompile(`\d+`)},
		{TokenIdentifier, regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`)},
		{TokenOperator, regexp.MustCompile(`[+\-*/]`)},
		{TokenAssign, regexp.MustCompile(`:=`)},
	}

	i := 0
	for i < len(input) {
		// Skip whitespace
		if strings.TrimSpace(string(input[i])) == "" {
			i++
			continue
		}

		// Try matching each token regex
		matched := false
		for _, spec := range tokenSpec {
			loc := spec.regx.FindStringIndex(input[i:])
			if loc != nil && loc[0] == 0 {
				tokens = append(tokens, Token{Type: spec.typ, Value: input[i : i+loc[1]]})
				i += loc[1]
				matched = true
				break
			}
		}

		if !matched {
			fmt.Println("Tokenization Error: Unrecognized symbol")
			tokens = append(tokens, Token{Type: TokenError, Value: string(input[i])})
			break
		}
	}

	// Append end of file token
	tokens = append(tokens, Token{Type: TokenEOF, Value: ""})
	return tokens
}
