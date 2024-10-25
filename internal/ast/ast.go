package ast

import (
	"fmt"
	"strconv"

	"github.com/algoboyz/argo/internal/tokenizer"
)

type ASTNode interface{}

// Different kinds of AST nodes
type NumberNode struct {
	Value int
}

type VarNode struct {
	Name string
}

type BinOpNode struct {
	Left     ASTNode
	Operator string
	Right    ASTNode
}

type AssignmentNode struct {
	Name  string
	Value ASTNode
}

type Parser struct {
	Tokens []tokenizer.Token
	pos    int
}

func (p *Parser) currentToken() tokenizer.Token {
	if p.pos < len(p.Tokens) {
		return p.Tokens[p.pos]
	}
	return tokenizer.Token{Type: tokenizer.TokenEOF}
}

func (p *Parser) eat(tokenType tokenizer.TokenType) tokenizer.Token {
	token := p.currentToken()
	if token.Type == tokenType {
		p.pos++
		return token
	}
	panic(fmt.Sprintf("Unexpected token: expected %s, got %s", tokenType, token.Type))
}

func (p *Parser) parseExpression() ASTNode {
	left := p.parseTerm()

	for p.currentToken().Type == tokenizer.TokenOperator {
		op := p.eat(tokenizer.TokenOperator).Value
		right := p.parseTerm()
		left = &BinOpNode{Left: left, Operator: op, Right: right}
	}

	return left
}

func (p *Parser) parseTerm() ASTNode {
	token := p.currentToken()

	switch token.Type {
	case tokenizer.TokenNumber:
		p.eat(tokenizer.TokenNumber)
		value, _ := strconv.Atoi(token.Value)
		return &NumberNode{Value: value}
	case tokenizer.TokenIdentifier:
		p.eat(tokenizer.TokenIdentifier)
		return &VarNode{Name: token.Value}
	default:
		panic("Unexpected token in term")
	}
}

func (p *Parser) parseAssignment() ASTNode {
	name := p.eat(tokenizer.TokenIdentifier).Value
	p.eat(tokenizer.TokenAssign)
	value := p.parseExpression()
	return &AssignmentNode{Name: name, Value: value}
}

func (p *Parser) Parse() ASTNode {
	return p.parseAssignment()
}
