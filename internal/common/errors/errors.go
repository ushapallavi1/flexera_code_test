package errors

import err "errors"

var ErrNotACSVRow = err.New("invalid csv data row")
var ErrOpenFile = err.New("error occurred while opening file")
var ErrReadingCsvDataRow = err.New("error occurred while reading csv data set row")
var ErrInvalidDataRow = err.New("received invalid data set row after parsing")
var ErrParsingCSVDataRow = err.New("error occurred while parsing csv data row")