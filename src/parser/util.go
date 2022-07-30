package main

import "go/ast"

// returns "" if failure
func exprToIdentName(expr ast.Expr) string {
	if id, ok := expr.(*ast.Ident); ok {
		return id.Name
	}
	if star, ok := expr.(*ast.StarExpr); ok {
		return exprToIdentName(star.X)
	}
	return ""
}
