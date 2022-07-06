package main

type parseResult struct {
	status             parseStatus
	failureMessage     string
	errorCodeLocations []*location
}

type parseStatus string

const (
	success parseStatus = "success"
	failure parseStatus = "failure"
)

type location struct {
	startLine int
	endLine   int
}

func newParseResult() *parseResult {
	return &parseResult{
		status:             success,
		errorCodeLocations: []*location{},
	}
}
