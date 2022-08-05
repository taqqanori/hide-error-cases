package main

type parseResult struct {
	Status             parseStatus `json:"status"`
	FailureMessage     string      `json:"failureMessage"`
	ErrorCodeLocations []*location `json:"errorCodeLocations"`
}

type parseStatus string

const (
	success parseStatus = "success"
	failure parseStatus = "failure"
)

type location struct {
	Start          *position `json:"start"`
	End            *position `json:"end"`
	BlockStartLine int       `json:"blockStartLine"`
}

type position struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

func newParseResult() *parseResult {
	return &parseResult{
		Status:             success,
		ErrorCodeLocations: []*location{},
	}
}
