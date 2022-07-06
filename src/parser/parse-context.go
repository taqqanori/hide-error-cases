package main

import (
	"go/ast"
	"go/token"
)

type parseContext struct {
	fileSet        *token.FileSet
	result         *parseResult
	currentDepth   int
	funcScopeStack *scopeStack
	ifScopeStack   *scopeStack
}

func (ctx *parseContext) Visit(node ast.Node) ast.Visitor {
	if ctx.result.status == failure {
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
		ctx.funcScopeStack.push(newFromFuncDecl(ctx.currentDepth, castedNode))
	case *ast.FuncLit:
		ctx.funcScopeStack.push(newFromFuncLit(ctx.currentDepth, castedNode))
	case *ast.IfStmt:
		ctx.ifScopeStack.push(newFromIfStmt(ctx.fileSet, ctx.currentDepth, castedNode))
	case *ast.ReturnStmt:
		ifScope, ok := ctx.ifScopeStack.peek().(*ifScope)
		if !ok {
			break
		}
		if ifScope == nil || ifScope.blockStartLine < 0 || ifScope.blockEndLine < 0 {
			// broken if statement
			break
		}
		funcScope, ok := ctx.funcScopeStack.peek().(*funcScope)
		if !ok {
			break
		}
		if len(castedNode.Results) <= funcScope.errorReturnTypeIndex {
			// return values and function return types does not match
			break
		}
		if exprToIdentName(castedNode.Results[funcScope.errorReturnTypeIndex]) == "nil" {
			// returning nil for error type, not a error case
			break
		}
		ctx.result.errorCodeLocations = append(ctx.result.errorCodeLocations, &location{
			startLine: ifScope.blockStartLine,
			endLine:   ifScope.blockEndLine,
		})
	}
	return ctx
}

func newParseContext(fset *token.FileSet, result *parseResult) *parseContext {
	return &parseContext{
		fileSet:        fset,
		result:         result,
		currentDepth:   0,
		funcScopeStack: &scopeStack{},
		ifScopeStack:   &scopeStack{},
	}
}
