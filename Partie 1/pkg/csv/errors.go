package csv

import "errors"

var (
	errNoPathGiven           = errors.New("no path given")
	errInvalidRecord         = errors.New("invalid record")
	errUnsupportedDateLayout = errors.New("unsupported date layout")
)
