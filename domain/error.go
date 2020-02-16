package domain

// Error エラーを扱うドメイン
type Error struct {
	code int
	err  error
}

func NewError(code int, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

func (this *Error) Code() int {
	return this.code
}

func (this *Error) Err() error {
	return this.err
}
