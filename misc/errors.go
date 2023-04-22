package misc

import "errors"

var (
	ErrNotFound       = errors.New("entity not found")
	ErrUserCtxExtract = errors.New("can not extract user context in schedule message handler")
)
