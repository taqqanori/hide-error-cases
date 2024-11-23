package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"regexp"
)

const defaultErrotTypeRegexp = "(E|e)rror$"

func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		errorOut("faild to read stdin.")
		return
	}

	errorTypeRegexpStr := defaultErrotTypeRegexp
	if 1 < len(os.Args) {
		errorTypeRegexpStr = os.Args[1]
	}

	json, err := json.Marshal(parse(string(src), errorTypeRegexpStr))
	if err != nil {
		errorOut("failed in json marshaling.")
		return
	}
	fmt.Print(string(json))
}

func errorOut(msg string) {
	fmt.Printf(`{"status":"failure","failureMessage":"%s"}`, msg)
}

func parse(src string, errorTypeRegexpStr string) *parseResult {
	ret := newParseResult()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		ret.Status = failure
		ret.FailureMessage = "Failed to parse file."
		return ret
	}

	errorTypeRegexp, err := regexp.Compile(errorTypeRegexpStr)
	if err != nil {
		errorTypeRegexp = regexp.MustCompile(defaultErrotTypeRegexp)
	}

	ctx := newParseContext(fset, ret, errorTypeRegexp)
	ast.Walk(ctx, f)

	return ret
}
