package encode

const urlLimit = 6

type URLEncoder interface {
	Encode([]byte) string
}

type URL struct {
	baseToEncode URLEncoder
}

func NewEncode(encode URLEncoder) URL {
	return URL{baseToEncode: encode}
}

func (u *URL) Encode(url []byte) string {
	return u.baseToEncode.Encode(url)[:urlLimit]
}
