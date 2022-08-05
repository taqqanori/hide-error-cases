package main

import (
	"go/ast"
	"go/token"
)

// covers else-if and else block
type ifScope struct {
	depth          int
	start          *position // nil if AST invalid
	end            *position // nil if AST invalid
	blockStartLine int
	ifStmt         *ast.IfStmt
}

func (is *ifScope) getDepth() int {
	return is.depth
}

func newFromIfStmt(fset *token.FileSet, depth int, ifStmt *ast.IfStmt) *ifScope {
	ret := &ifScope{
		depth:  depth,
		ifStmt: ifStmt,
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
	ret.blockStartLine = start.Line
	ret.end = &position{
		Line:   end.Line,
		Column: end.Column,
	}
	return ret
}

func newFromElseStmt(fset *token.FileSet, depth int, ifStmt *ast.IfStmt) *ifScope {
	// AST does not tell the "else" position...
	// But fortunately, "}" of previous if block and "else" must be in same line, so let (position of "}") + 1 be the start of "else".
	// https://stackoverflow.com/questions/26371645/unexpected-semicolon-or-newline-before-else-even-though-there-is-neither-before/26371912#26371912
	ff := fset.File(ifStmt.Else.Pos())
	start := ff.Position(ifStmt.Body.Rbrace + 1)

	switch elseStmt := ifStmt.Else.(type) {
	case *ast.IfStmt:
		// else if {...}
		ret := newFromIfStmt(fset, depth, elseStmt)
		// just reset the start position to "else"
		ret.start = &position{
			Line:   start.Line,
			Column: start.Column,
		}
		return ret
	case *ast.BlockStmt:
		// else {...}
		ret := &ifScope{
			depth: depth,
			start: &position{
				Line:   start.Line,
				Column: start.Column,
			},
		}
		blockStart := ff.Position(elseStmt.Lbrace)
		ret.blockStartLine = blockStart.Line
		end := ff.Position(elseStmt.Rbrace)
		ret.end = &position{
			Line:   end.Line,
			Column: end.Column,
		}
		return ret
	default:
		// not sure about this case, maybe compiler has gone mad?
		return &ifScope{
			depth: depth,
		}
	}
}
