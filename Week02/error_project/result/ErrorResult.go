package result

type ErrorResult struct {
	err error
}

func NewErrorResult(err error) *ErrorResult {
	return &ErrorResult{err: err}
}



