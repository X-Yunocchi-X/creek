package creek

import (
	"github.com/X-Yunocchi-X/creek/judge"
)

// TODO: unsigned int, overflow check, other for self-function
func Range[T judge.SignedInt](start, end T, other ...any) []T {
	res := make([]T, end-start)
	for i := range res {
		res[i] = start + T(i)
	}
	return res
}

// TODO: { x | x ∈ setA ∧ x ∈ setB }

// TODO: pattern match
