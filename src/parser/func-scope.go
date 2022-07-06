package main

import "go/ast"

type funcScope struct {
	depth                int
	errorReturnTypeIndex int // -1 if it does not return error
}

func (fc *funcScope) getDepth() int {
	return fc.depth
}

func newFromFuncDecl(depth int, decl *ast.FuncDecl) *funcScope {
	return newFromFuncType(depth, decl.Type)
}

func newFromFuncLit(depth int, lit *ast.FuncLit) *funcScope {
	return newFromFuncType(depth, lit.Type)
}

func newFromFuncType(depth int, funcType *ast.FuncType) *funcScope {
	ret := &funcScope{
		depth:                depth,
		errorReturnTypeIndex: -1,
	}
	if funcType == nil || funcType.Results == nil || funcType.Results.List == nil {
		return ret
	}
	for i := 0; i < len(funcType.Results.List); i++ {
		if exprToIdentName(funcType.Results.List[i].Type) == "error" {
			ret.errorReturnTypeIndex = i
		}
	}
	return ret
}
