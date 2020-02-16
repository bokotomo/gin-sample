package domain

// Error エラーを扱うドメイン
type Error struct {
	code int
	err  error
}

// NewError is
func NewError(code int, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

// Code is
func (e *Error) Code() int {
	return e.code
}

// Err is
func (e *Error) Err() error {
	return e.err
}
