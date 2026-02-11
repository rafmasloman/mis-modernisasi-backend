package errors

import "errors"

var (
	ErrReportNotFound     = errors.New("report not found")
	ErrReportInactive     = errors.New("report is inactive")
	ErrInvalidReportQuery = errors.New("invalid report query")
)
