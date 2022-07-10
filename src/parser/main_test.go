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
		Status: success,
		ErrorCodeLocations: []*location{
			{StartLine: 9, EndLine: 12},
			{StartLine: 26, EndLine: 29},
		},
	})
	test(t, "literals.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			{StartLine: 8, EndLine: 10},
			{StartLine: 17, EndLine: 19},
		},
	})
	test(t, "named-return-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			{StartLine: 7, EndLine: 9},
			{StartLine: 14, EndLine: 16},
			{StartLine: 23, EndLine: 25},
		},
	})
	test(t, "loop.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			{
				StartLine: 8,
				EndLine:   11,
			},
		},
	})
	test(t, "multiple-errors.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			{StartLine: 6, EndLine: 8},
			{StartLine: 10, EndLine: 12},
			{StartLine: 14, EndLine: 16},
			{StartLine: 18, EndLine: 20},
			{StartLine: 22, EndLine: 24},
		},
	})
	test(t, "nested.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			{StartLine: 7, EndLine: 9},
			{StartLine: 6, EndLine: 11},
			{StartLine: 16, EndLine: 18},
		},
	})
	test(t, "if-outside-func.go", parseResult{
		Status:             success,
		ErrorCodeLocations: []*location{},
	})
	parseAndCompare(t, "I am invalid as go code!", parseResult{
		Status:             failure,
		FailureMessage:     "Failed to parse file.",
		ErrorCodeLocations: []*location{},
	})
}

func test(t *testing.T, file string, expected parseResult) {
	_, self, _, _ := runtime.Caller(0)
	src, _ := ioutil.ReadFile(filepath.Join(filepath.Dir(self), testInputsDir, file))
	parseAndCompare(t, string(src), expected)
}

func parseAndCompare(t *testing.T, src string, expected parseResult) {
	assert.JSONEq(t, j(t, expected), j(t, parse(src)))
}

func j(t *testing.T, o interface{}) string {
	b, err := json.Marshal(o)
	assert.Nil(t, err)
	return string(b)
}
