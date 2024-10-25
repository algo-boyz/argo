package asm

import (
	"fmt"

	"github.com/algoboyz/argo/internal/ast"
)

type CodeGenerator struct {
	registerCount int
}

func (cg *CodeGenerator) NewRegister() string {
	reg := fmt.Sprintf("X%d", cg.registerCount)
	cg.registerCount++
	return reg
}

func (cg *CodeGenerator) Visit(node ast.ASTNode) string {
	switch n := node.(type) {
	case *ast.NumberNode:
		reg := cg.NewRegister()
		return fmt.Sprintf("MOV %s, #%d\n", reg, n.Value)
	case *ast.VarNode:
		// for full implementation, handle variable lookup here
		return fmt.Sprintf("MOV %s, [VAR]\n", n.Name)
	case *ast.BinOpNode:
		leftCode := cg.Visit(n.Left)
		rightCode := cg.Visit(n.Right)
		reg := cg.NewRegister()
		return fmt.Sprintf("%s%sADD %s, X0, X1\n", leftCode, rightCode, reg)
	case *ast.AssignmentNode:
		valueCode := cg.Visit(n.Value)
		return fmt.Sprintf("%sSTR X0, [%s]\n", valueCode, n.Name)
	}
	return ""
}
