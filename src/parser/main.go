package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		errorOut("faild to read stdin.")
	}
	json, err := json.Marshal(parse(string(src)))
	if err != nil {
		errorOut("failed in json marshaling.")
		return
	}
	fmt.Print(string(json))
}

func errorOut(msg string) {
	fmt.Printf(`{status:"failure",failureMessage:"%s"}`, msg)
}

func parse(src string) *parseResult {
	ret := newParseResult()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		ret.Status = failure
		ret.FailureMessage = "Failed to parse file."
		return ret
	}

	ctx := newParseContext(fset, ret)
	ast.Walk(ctx, f)

	return ret
}
