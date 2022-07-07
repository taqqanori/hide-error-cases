package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testInputsDir = "test-inputs"

func Test(t *testing.T) {
	test(t, "general.go", parseResult{
		status:             success,
		errorCodeLocations: []*location{{startLine: 9, endLine: 12}},
	})
	test(t, "literals.go", parseResult{
		status: success,
		errorCodeLocations: []*location{
			{startLine: 8, endLine: 10},
			{startLine: 17, endLine: 19},
		},
	})
	test(t, "named-return-types.go", parseResult{
		status: success,
		errorCodeLocations: []*location{
			{startLine: 7, endLine: 9},
			{startLine: 14, endLine: 16},
			{startLine: 23, endLine: 25},
		},
	})
	test(t, "loop.go", parseResult{
		status: success,
		errorCodeLocations: []*location{
			{
				startLine: 8,
				endLine:   11,
			},
		},
	})
	test(t, "multiple-errors.go", parseResult{
		status: success,
		errorCodeLocations: []*location{
			{startLine: 6, endLine: 8},
			{startLine: 10, endLine: 12},
			{startLine: 14, endLine: 16},
			{startLine: 18, endLine: 20},
			{startLine: 22, endLine: 24},
		},
	})
}

func test(t *testing.T, file string, expected parseResult) {
	_, self, _, _ := runtime.Caller(0)
	content, _ := ioutil.ReadFile(filepath.Join(filepath.Dir(self), testInputsDir, file))
	assert.JSONEq(t, j(t, expected), j(t, parse(string(content))))
}

func j(t *testing.T, o interface{}) string {
	b, err := json.Marshal(o)
	assert.Nil(t, err)
	return string(b)
}
