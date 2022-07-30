package main

import (
	"go/ast"
	"go/token"
	"regexp"
)

type parseContext struct {
	fileSet         *token.FileSet
	result          *parseResult
	errorTypeRegexp *regexp.Regexp
	currentDepth    int
	funcScopeStack  *scopeStack
	ifScopeStack    *scopeStack
}

func (ctx *parseContext) Visit(node ast.Node) ast.Visitor {
	if ctx.result.Status == failure {
		return nil
	}
	if node == nil {
		ctx.funcScopeStack.popIfDepthMatch(ctx.currentDepth)
		ctx.ifScopeStack.popIfDepthMatch(ctx.currentDepth)
		ctx.currentDepth--
		return nil
	}
	ctx.currentDepth++
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
		for _, errorReturnTypeIndex := range funcScope.errorReturnTypeIndices {
			if len(castedNode.Results) <= errorReturnTypeIndex {
				// return values and function return types does not match
				continue
			}
			if exprToIdentName(castedNode.Results[errorReturnTypeIndex]) == "nil" {
				// returning nil for error type, not a error case
				continue
			}
			ctx.result.ErrorCodeLocations = append(ctx.result.ErrorCodeLocations, &location{
				Start: ifScope.start,
				End:   ifScope.end,
			})
			break
		}
	}
	return ctx
}

func newParseContext(fset *token.FileSet, result *parseResult, errorTypeRegexp *regexp.Regexp) *parseContext {
	return &parseContext{
		fileSet:         fset,
		result:          result,
		errorTypeRegexp: errorTypeRegexp,
		currentDepth:    0,
		funcScopeStack:  &scopeStack{},
		ifScopeStack:    &scopeStack{},
	}
}
