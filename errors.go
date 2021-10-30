package go_pool

import "errors"

var (
	ErrIncorrectSize = errors.New("incorrect pool size: should be >= 1 and <= 65535")
)
