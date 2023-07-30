package slice

type Signed interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8
}

type Unsigned interface {
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Float | Integer
}
