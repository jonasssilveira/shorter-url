package encode

const urlLimit = 6

type EncodeURL interface {
	Encode([]byte) string
}

type URL struct {
	baseToEncode EncodeURL
}

func NewEncode(encode EncodeURL) URL {
	return URL{baseToEncode: encode}
}

func (u *URL) Encode(url string) string {
	return u.baseToEncode.Encode([]byte(url))[:urlLimit]
}
