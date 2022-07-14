package main

import (
	"go/ast"
	"go/token"
)

type ifScope struct {
	depth int
	start *position // nil if AST invalid
	end   *position // nil1 if AST invalid
}

func (is *ifScope) getDepth() int {
	return is.depth
}

func newFromIfStmt(fset *token.FileSet, depth int, ifStmt *ast.IfStmt) *ifScope {
	ret := &ifScope{
		depth: depth,
	}
	if ifStmt.Body == nil {
		return ret
	}
	ff := fset.File(ifStmt.Pos())
	start := ff.Position(ifStmt.If)
	end := ff.Position(ifStmt.Body.Rbrace)
	ret.start = &position{
		Line:   start.Line,
		Column: start.Column,
	}
	ret.end = &position{
		Line:   end.Line,
		Column: end.Column,
	}
	return ret
}
