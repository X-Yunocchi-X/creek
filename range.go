package creek

import (
	"github.com/X-Yunocchi-X/creek/judge"
)

// TODO: unsigned int, overflow check, other for self-function
func Range[T judge.SignedInt](start, end T, other ...func(T) T) []T {
	if len(other) > 1 {
		return nil
	}
	res := make([]T, end-start)
	n := len(res)
	f := func(t T) T {
		return t
	}
	if len(other) == 1 {
		f = other[0]
	}
	for i := 0; i < n; i++ {
		res[i] = f(T(i) + start)
	}
	return res
}

// TODO: { x | x ∈ setA ∧ x ∈ setB }

// TODO: pattern match
