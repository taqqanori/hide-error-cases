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
	StartLine int `json:"startLine"`
	EndLine   int `json:"endLine"`
}

func newParseResult() *parseResult {
	return &parseResult{
		Status:             success,
		ErrorCodeLocations: []*location{},
	}
}
