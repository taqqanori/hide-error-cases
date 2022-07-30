package main

import (
	"go/ast"
	"regexp"
)

type funcScope struct {
	depth                  int
	errorReturnTypeIndices []int // empty if it does not return error
}

func (fc *funcScope) getDepth() int {
	return fc.depth
}

func newFromFuncDecl(depth int, decl *ast.FuncDecl, errorTypeRegexp *regexp.Regexp) *funcScope {
	return newFromFuncType(depth, decl.Type, errorTypeRegexp)
}

func newFromFuncLit(depth int, lit *ast.FuncLit, errorTypeRegexp *regexp.Regexp) *funcScope {
	return newFromFuncType(depth, lit.Type, errorTypeRegexp)
}

func newFromFuncType(depth int, funcType *ast.FuncType, errorTypeRegexp *regexp.Regexp) *funcScope {
	ret := &funcScope{
		depth:                  depth,
		errorReturnTypeIndices: []int{},
	}
	if funcType == nil || funcType.Results == nil || funcType.Results.List == nil {
		return ret
	}
	for i := 0; i < len(funcType.Results.List); i++ {
		if errorTypeRegexp.MatchString(exprToIdentName(funcType.Results.List[i].Type)) {
			ret.errorReturnTypeIndices = append(ret.errorReturnTypeIndices, i)
		}
	}
	return ret
}
