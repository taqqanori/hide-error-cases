package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"regexp"
)

const defaultErrorTypeRegexp = "(E|e)rror$"

func main() {
	src, err := io.ReadAll(os.Stdin)
	if err != nil {
		errorOut("Failed to read stdin.")
		return
	}

	errorTypeRegexpStr := defaultErrorTypeRegexp
	if 1 < len(os.Args) {
		errorTypeRegexpStr = os.Args[1]
	}

	json, err := json.Marshal(parse(string(src), errorTypeRegexpStr))
	if err != nil {
		errorOut("Failed in json marshaling.")
		return
	}
	fmt.Print(string(json))
}

func errorOut(msg string) {
	fmt.Printf(`{"status":"failure","failureMessage":"%s"}`, msg)
}

func parse(src string, errorTypeRegexpStr string) *parseResult {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		ret := newParseResult()
		ret.Status = failure
		ret.FailureMessage = "Failed to parse file."
		return ret
	}

	errorTypeRegexp, err := regexp.Compile(errorTypeRegexpStr)
	if err != nil {
		errorTypeRegexp = regexp.MustCompile(defaultErrorTypeRegexp)
	}

	ctx := newParseContext(fset, errorTypeRegexp)
	ast.Walk(ctx, f)
	ret := ctx.Result()

	return ret
}
