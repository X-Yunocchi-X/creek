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

func Multi[T Judgable](v1, v2 T) T {
	return v1 * v2
}

// TODO: Exp function
func Exp[T Judgable](v T, x int) T {
	return exp(v, x)
}

// ! x >= 0, v >= 0
func exp[T Judgable](v T, x int) T {
	if v == 0 && x == 0 {
		panic("0^0 is undefined")
	}
	if v == T(0) {
		return T(0)
	}
	if x == 0 {
		return T(1)
	}
	if x < 0 || v < 0 {
		panic("shouldn't acheive this")
	}
	// TODO: fast pow algorithm
	var res T
	for i := 1; i <= x; i++ {
		res *= v
	}
	return res
}
