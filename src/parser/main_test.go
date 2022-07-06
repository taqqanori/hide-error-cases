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

func TestHoge(t *testing.T) {
	test(t, "general.go", parseResult{
		status: success,
		errorCodeLocations: []*location{
			{
				startLine: 7,
				endLine:   9,
			},
			{
				startLine: 14,
				endLine:   16,
			},
			{
				startLine: 23,
				endLine:   25,
			},
			{
				startLine: 38,
				endLine:   40,
			},
			{
				startLine: 45,
				endLine:   47,
			},
			{
				startLine: 54,
				endLine:   56,
			},
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
