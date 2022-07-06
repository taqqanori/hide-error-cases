package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fmt.Println("hoge")
}

func parse(src string) *parseResult {
	ret := newParseResult()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		ret.status = failure
		ret.failureMessage = "Failed to parse file."
		return ret
	}

	ctx := newParseContext(fset, ret)
	ast.Walk(ctx, f)

	return ret
}
