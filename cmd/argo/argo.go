package main

import (
	"fmt"

	"github.com/algoboyz/argo/internal/asm"
	"github.com/algoboyz/argo/internal/ast"
	"github.com/algoboyz/argo/internal/tokenizer"
)

func main() {
	code := "x := 42 + 5"

	tokens := tokenizer.Tokenize(code)
	for _, token := range tokens {
		fmt.Printf("Token: %+v\n", token)
	}

	parser := &ast.Parser{Tokens: tokens}
	ast := parser.Parse()
	fmt.Printf("AST: %+v\n", ast)

	codegen := &asm.CodeGenerator{}
	assembly := codegen.Visit(ast)
	fmt.Println(assembly)
}
