package main

import (
	"go/ast"
	"go/token"
	"regexp"
)

type parseContext struct {
	fileSet           *token.FileSet
	errorTypeRegexp   *regexp.Regexp
	currentDepth      int
	funcScopeStack    *scopeStack
	ifScopeStack      *scopeStack
	candidateIfScopes []*ifScope
	whiteListIfScopes []*ifScope
}

func (ctx *parseContext) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		ctx.funcScopeStack.popIfDepthMatch(ctx.currentDepth)
		ctx.ifScopeStack.popIfDepthMatch(ctx.currentDepth)
		ctx.currentDepth--
		return nil
	}
	ctx.currentDepth++
	if ifs, ok := ctx.ifScopeStack.peek().(*ifScope); ok && ifs.ifStmt != nil && node == ifs.ifStmt.Else {
		// now came down to "else" ("else" is located under "if" on AST)
		ctx.ifScopeStack.push(newFromElseStmt(ctx.fileSet, ctx.currentDepth, ifs.ifStmt))
		return ctx
	}
	switch castedNode := node.(type) {
	case *ast.FuncDecl:
		ctx.funcScopeStack.push(newFromFuncDecl(ctx.currentDepth, castedNode, ctx.errorTypeRegexp))
	case *ast.FuncLit:
		ctx.funcScopeStack.push(newFromFuncLit(ctx.currentDepth, castedNode, ctx.errorTypeRegexp))
	case *ast.IfStmt:
		ctx.ifScopeStack.push(newFromIfStmt(ctx.fileSet, ctx.currentDepth, castedNode))
	case *ast.ReturnStmt:
		ifScope, ok := ctx.ifScopeStack.peek().(*ifScope)
		if !ok {
			break
		}
		if ifScope == nil || ifScope.start == nil || ifScope.end == nil {
			// broken if statement
			break
		}
		funcScope, ok := ctx.funcScopeStack.peek().(*funcScope)
		if !ok {
			break
		}
		if ifScope.depth <= funcScope.depth {
			// this if statement is outside of the func
			break
		}
		returningError := false
		if len(castedNode.Results) == 0 && 0 < len(funcScope.errorReturnTypes) {
			// named return values and return statement without arguments case
			// http://go.dev/tour/basics/7
			// https://github.com/taqqanori/hide-error-cases/issues/5
			returningError = isErrorCaseCond(ifScope.ifStmt.Cond, funcScope.errorReturnTypes)
		} else {
			for _, errorReturnType := range funcScope.errorReturnTypes {
				if len(castedNode.Results) <= errorReturnType.index {
					// return values and function return types does not match
					continue
				}
				if exprToIdentName(castedNode.Results[errorReturnType.index]) == "nil" {
					// returning nil for error type
					continue
				}
				// returning something not nil for error type
				returningError = true
				break
			}
		}
		if returningError {
			// returning error for at least one error type, add to candidate list
			ctx.candidateIfScopes = append(ctx.candidateIfScopes, ifScope)
		} else {
			// returning nil for all the error types, not a error case, add to white list,
			// for the case ifScope has multiple return statements
			// https://github.com/taqqanori/hide-error-cases/issues/6
			ctx.whiteListIfScopes = append(ctx.whiteListIfScopes, ifScope)
		}
	}
	return ctx
}

func (ctx *parseContext) Result() *parseResult {
	ret := newParseResult()
	for _, ifScope := range ctx.candidateIfScopes {
		whiteListed := false
		for _, white := range ctx.whiteListIfScopes {
			if ifScope == white {
				whiteListed = true
			}
		}
		if !whiteListed {
			ret.ErrorCodeLocations = append(ret.ErrorCodeLocations, &location{
				Start:          ifScope.start,
				End:            ifScope.end,
				BlockStartLine: ifScope.blockStartLine,
			})
		}
	}
	return ret
}

func isErrorCaseCond(expr ast.Expr, errorReturnTypes []*errorReturnType) bool {
	if parenExpr, ok := expr.(*ast.ParenExpr); ok {
		// ( X )
		return isErrorCaseCond(parenExpr.X, errorReturnTypes)
	}
	if binaryExpr, ok := expr.(*ast.BinaryExpr); ok {
		switch binaryExpr.Op {
		case token.LAND:
			return isErrorCaseCond(binaryExpr.X, errorReturnTypes) || isErrorCaseCond(binaryExpr.Y, errorReturnTypes)
		case token.LOR:
			return isErrorCaseCond(binaryExpr.X, errorReturnTypes) && isErrorCaseCond(binaryExpr.Y, errorReturnTypes)
		case token.NEQ:
			for _, errorReturnType := range errorReturnTypes {
				if exprToIdentName(binaryExpr.X) == errorReturnType.name && exprToIdentName(binaryExpr.Y) == "nil" {
					return true
				}
				if exprToIdentName(binaryExpr.Y) == errorReturnType.name && exprToIdentName(binaryExpr.X) == "nil" {
					return true
				}
			}
		}
	}
	return false
}

func newParseContext(fset *token.FileSet, errorTypeRegexp *regexp.Regexp) *parseContext {
	return &parseContext{
		fileSet:           fset,
		errorTypeRegexp:   errorTypeRegexp,
		currentDepth:      0,
		funcScopeStack:    &scopeStack{},
		ifScopeStack:      &scopeStack{},
		candidateIfScopes: []*ifScope{},
		whiteListIfScopes: []*ifScope{},
	}
}
