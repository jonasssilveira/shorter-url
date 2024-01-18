package usecase

type EncodeAdapter struct {
	fn func([]byte) string
}

func NewEncodeAdapter(fn func([]byte) string) EncodeAdapter {
	return EncodeAdapter{fn: fn}
}

func (e *EncodeAdapter) Encode(url []byte) string {
	return e.fn(url)
}
