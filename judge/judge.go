package judge

type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	SignedInt | UnsignedInt
}

type Judgable interface {
	Integer | Float
}

func IsOdd[T Integer](v T) bool {
	return v%2 != 0
}

func IsEven[T Integer](v T) bool {
	return v%2 == 0
}

func Sum[T Judgable](v1, v2 T) T {
	return v1 + v2
}

func Max[T Judgable](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Min[T Judgable](v1, v2 T) T {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Abs[T SignedInt | Float](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

// TODO: Exp function
