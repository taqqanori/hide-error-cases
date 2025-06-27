package main

import (
	"go/ast"
	"regexp"
)

type funcScope struct {
	depth            int
	errorReturnTypes []*errorReturnType // empty if it does not return error
}

type errorReturnType struct {
	index int
	name  string
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
		depth:            depth,
		errorReturnTypes: []*errorReturnType{},
	}
	if funcType == nil || funcType.Results == nil || funcType.Results.List == nil {
		return ret
	}
	offset := 0
	for i := 0; i < len(funcType.Results.List); i++ {
		field := funcType.Results.List[i]
		if 0 < len(field.Names) {
			// may have multiple named return values with single type
			for j, name := range field.Names {
				if j != 0 {
					offset++
				}
				if errorTypeRegexp.MatchString(exprToIdentName(field.Type)) {
					ret.errorReturnTypes = append(ret.errorReturnTypes, &errorReturnType{
						index: offset + i,
						name:  exprToIdentName(name),
					})
				}
			}
		} else {
			if errorTypeRegexp.MatchString(exprToIdentName(field.Type)) {
				ret.errorReturnTypes = append(ret.errorReturnTypes, &errorReturnType{
					index: offset + i,
				})
			}
		}
	}
	return ret
}
