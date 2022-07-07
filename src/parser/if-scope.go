package main

import (
	"go/ast"
	"go/token"
)

type ifScope struct {
	depth          int
	blockStartLine int // -1 if AST invalid
	blockEndLine   int // -1 if AST invalid
}

func (is *ifScope) getDepth() int {
	return is.depth
}

func newFromIfStmt(fset *token.FileSet, depth int, ifStmt *ast.IfStmt) *ifScope {
	ret := &ifScope{
		depth:          depth,
		blockStartLine: -1,
		blockEndLine:   -1,
	}
	if ifStmt.Body == nil {
		return ret
	}
	ff := fset.File(ifStmt.Pos())
	ret.blockStartLine = ff.Line(ifStmt.Body.Lbrace)
	ret.blockEndLine = ff.Line(ifStmt.Body.Rbrace)
	return ret
}
