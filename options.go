package gomapper

type withFuncOption[TSource, TDest any] struct {
	fn func(TSource, *TDest)
}

type withIgnoreOption[TDest any] struct {
	fn func(*TDest) []any
}

type options struct {
	Fns          []any
	IgnoreFns    []any
	IgnoreFields map[string]bool
}

type Option interface {
	apply(*options)
}

func (a withFuncOption[TSource, TDest]) apply(opts *options) {
	opts.Fns = append(opts.Fns, a.fn)
}

func (a withIgnoreOption[TDest]) apply(opts *options) {
	if opts.IgnoreFields == nil {
		opts.IgnoreFields = make(map[string]bool)
	}
	opts.IgnoreFns = append(opts.IgnoreFns, a.fn)
}

func WithFunc[TSource, TDest any](fn func(TSource, *TDest)) Option {
	return &withFuncOption[TSource, TDest]{fn: fn}
}

func WithIgnore[TDest any](fn func(*TDest) []any) Option {
	return &withIgnoreOption[TDest]{fn: fn}
}
