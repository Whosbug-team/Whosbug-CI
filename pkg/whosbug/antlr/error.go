package antlr

import "errors"

var (
	// ErrAlreadyExists release已存在，且latest-commit未更新
	ErrAlreadyExists = errors.New("The Project and Release already exist")
)
