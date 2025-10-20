package server

import "errors"

var (
	// ErrMissingDatabase indicates a caller to construct a Server without providing
	// a database to connect to.
	ErrMissingDatabase = errors.New("must provide a database")
)
