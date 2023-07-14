package creek

import (
	"fmt"
)

// TODO: append, delete fun, change comparable
type Stream[T comparable] struct {
	arr []T
	len int
}

func New[T comparable](s []T) Stream[T] {
	res := Stream[T]{}
	res.init(s)
	return res
}

func (s Stream[T]) Len() int {
	return s.len
}

func (s *Stream[T]) init(arr []T) {
	s.len = len(arr)
	s.arr = make([]T, s.len)
	copy(s.arr, arr)
}

func (s Stream[T]) ToString() string {
	return fmt.Sprintf("%v", s.arr)
}

func (s Stream[T]) Map(f func(T) T) Stream[T] {
	res := Stream[T]{}
	res.init(make([]T, len(s.arr)))
	for i, v := range s.arr {
		res.arr[i] = f(v)
	}
	return res
}

func (s Stream[T]) Filter(f func(T) bool) Stream[T] {
	res := Stream[T]{}
	for _, v := range s.arr {
		if f(v) {
			res.arr = append(res.arr, v)
		}
	}
	return res
}

func (s Stream[T]) Foldl(f func(T, T) T, init T) T {
	res := init
	for _, v := range s.arr {
		res = f(res, v)
	}
	return res
}

func (s Stream[T]) Foldr(f func(T, T) T, init T) T {
	res := init
	for i := len(s.arr) - 1; i >= 0; i-- {
		res = f(s.arr[i], res)
	}
	return res
}

func (s Stream[T]) Fold(f func(T, T) T) T {
	var t T
	return s.Foldl(f, t)
}

func (s Stream[T]) Register(f func(T)) {
	for _, v := range s.arr {
		f(v)
	}
}

func (s Stream[T]) Any(f func(T) bool) bool {
	for _, v := range s.arr {
		if f(v) {
			return true
		}
	}
	return false
}

func (s Stream[T]) All(f func(T) bool) bool {
	return !s.Any(func(v T) bool { return !f(v) })
}

func (s Stream[T]) Distinct() Stream[T] {
	res := Stream[T]{}
	set := map[T]bool{}
	for _, v := range s.arr {
		if _, ok := set[v]; !ok {
			set[v] = true
			res.arr = append(res.arr, v)
		}
	}
	return res
}
