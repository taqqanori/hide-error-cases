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

type parseResult struct {
	status             parseStatus
	failureMessage     string
	errorCodeLocations []location
}

type parseStatus string

const (
	success parseStatus = "success"
	failure parseStatus = "failure"
)

type location struct {
	startLine int
	endLine   int
}

func parse(src string) parseResult {
	ret := parseResult{
		status:             success,
		errorCodeLocations: []location{},
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		ret.status = failure
		ret.failureMessage = "Failed to parse file."
		return ret
	}
	ast.Print(fset, f)

	ctx := parseContext{
		result: &ret,
	}
	for decl := range f.Decls {
		parseRecursive(decl, &ctx)
	}

	return ret
}

type parseContext struct {
	result *parseResult
}

func parseRecursive(node interface{}, result *parseContext) {

}
