package result

type Result struct {
	Value any
	Error error
}

func Ok(value any) *Result {
	return &Result{Value: value}
}

func Err(err error) *Result {
	return &Result{Error: err}
}

func (r *Result) IsOk() bool {
	return r.Error == nil
}

func (r *Result) IsErr() bool {
	return r.Error != nil
}

func Match[T any](r *Result, ok func(T), err func(error)) {
	if r.IsOk() {
		ok(r.Value.(T))
	}
	err(r.Error)
}

func (r *Result) ValueOrZero() any {
	if r.IsOk() {
		return r.Value
	}
	return nil
}

func (r *Result) ValueOrError() (any, error) {
	if r.IsOk() {
		return r.Value, nil
	}
	return nil, r.Error
}

func (r *Result) ValueOrPanic() any {
	if r.IsOk() {
		return r.Value
	}
	panic(r.Error)
}
