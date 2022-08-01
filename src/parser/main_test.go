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
			loc(9, 2, 12, 2),
			loc(26, 2, 29, 2),
		},
	})
	test(t, "literals.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(8, 3, 10, 3),
			loc(17, 3, 19, 3),
		},
	})
	test(t, "named-return-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2),
			loc(14, 3, 16, 3),
			loc(23, 3, 25, 3),
		},
	})
	test(t, "loop.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(8, 3, 11, 3),
		},
	})
	test(t, "multiple-errors.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(6, 2, 8, 2),
			loc(10, 2, 12, 2),
			loc(14, 2, 16, 2),
			loc(18, 2, 20, 2),
			loc(22, 2, 24, 2),
		},
	})
	test(t, "nested.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 3, 9, 3),
			loc(6, 2, 11, 2),
			loc(16, 4, 18, 4),
		},
	})
	test(t, "if-outside-func.go", parseResult{
		Status:             success,
		ErrorCodeLocations: []*location{},
	})
	test(t, "custom-error-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2),
			loc(26, 2, 28, 2),
		},
	})
	test(t, "custom-error-types-regexp.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2),
			loc(26, 2, 28, 2),
		},
	},
		"Exception")
	test(t, "custom-packaged-error-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2),
			loc(24, 2, 26, 2),
		},
	})
	parseAndCompare(t, "I am invalid as go code!", parseResult{
		Status:             failure,
		FailureMessage:     "Failed to parse file.",
		ErrorCodeLocations: []*location{},
	})
}

func loc(startLine int, startColumn int, endLine int, endColumn int) *location {
	return &location{
		Start: &position{
			Line:   startLine,
			Column: startColumn,
		},
		End: &position{
			Line:   endLine,
			Column: endColumn,
		},
	}
}

func test(t *testing.T, file string, expected parseResult, errorTypeRegex ...string) {
	_, self, _, _ := runtime.Caller(0)
	src, _ := ioutil.ReadFile(filepath.Join(filepath.Dir(self), testInputsDir, file))
	parseAndCompare(t, string(src), expected, errorTypeRegex...)
}

func parseAndCompare(t *testing.T, src string, expected parseResult, errorTypeRegexArr ...string) {
	errorTypeRegex := defaultErrotTypeRegexp
	if 0 < len(errorTypeRegexArr) {
		errorTypeRegex = errorTypeRegexArr[0]
	}
	assert.JSONEq(t, j(t, expected), j(t, parse(src, errorTypeRegex)))
}

func j(t *testing.T, o interface{}) string {
	b, err := json.Marshal(o)
	assert.Nil(t, err)
	return string(b)
}
