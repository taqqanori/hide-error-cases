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
			loc(9, 2, 12, 2, 9),
			loc(26, 2, 29, 2, 26),
		},
	})
	test(t, "literals.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(8, 3, 10, 3, 8),
			loc(17, 3, 19, 3, 17),
		},
	})
	test(t, "named-return-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2, 7),
			loc(14, 3, 16, 3, 14),
			loc(23, 3, 25, 3, 23),
		},
	})
	test(t, "loop.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(8, 3, 11, 3, 8),
		},
	})
	test(t, "multiple-errors.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(6, 2, 8, 2, 6),
			loc(10, 2, 12, 2, 10),
			loc(14, 2, 16, 2, 14),
			loc(18, 2, 20, 2, 18),
			loc(22, 2, 24, 2, 22),
		},
	})
	test(t, "nested.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 3, 9, 3, 7),
			loc(6, 2, 11, 2, 6),
			loc(16, 4, 18, 4, 16),
		},
	})
	test(t, "if-outside-func.go", parseResult{
		Status:             success,
		ErrorCodeLocations: []*location{},
	})
	test(t, "custom-error-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2, 7),
			loc(26, 2, 28, 2, 26),
		},
	})
	test(t, "custom-error-types-regexp.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2, 7),
			loc(26, 2, 28, 2, 26),
		},
	},
		"Exception")
	test(t, "custom-packaged-error-types.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(7, 2, 9, 2, 7),
			loc(24, 2, 26, 2, 24),
		},
	})
	parseAndCompare(t, "I am invalid as go code!", parseResult{
		Status:             failure,
		FailureMessage:     "Failed to parse file.",
		ErrorCodeLocations: []*location{},
	})
}

func TestXxx(t *testing.T) {
	test(t, "if-else.go", parseResult{
		Status: success,
		ErrorCodeLocations: []*location{
			loc(9, 3, 12, 2, 9),
			loc(17, 2, 20, 2, 17),
			loc(29, 3, 32, 2, 29),
			loc(38, 2, 41, 2, 38),
			loc(49, 2, 52, 2, 49),
			loc(63, 3, 66, 2, 63),
			loc(77, 3, 80, 2, 77),
			loc(88, 4, 90, 3, 88),
			loc(90, 4, 92, 3, 90),
			loc(85, 2, 94, 2, 85),
			loc(95, 3, 97, 3, 95),
			loc(99, 4, 101, 3, 99),
			loc(104, 3, 106, 3, 104),
			loc(106, 4, 108, 3, 106),
			loc(103, 3, 112, 2, 103),
			loc(119, 3, 123, 2, 120),
		},
	})
}

func loc(startLine int, startColumn int, endLine int, endColumn int, blockStartLine int) *location {
	return &location{
		Start: &position{
			Line:   startLine,
			Column: startColumn,
		},
		End: &position{
			Line:   endLine,
			Column: endColumn,
		},
		BlockStartLine: blockStartLine,
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
