package Slice

func check(c, l int) (int, bool) {
	if c <= 64 {
		return c, false
	} else if c > 2048 && c/l >= 2 {
		return int(float64(c) * 0.625), true
	} else if c <= 2048 && c/l >= 4 {
		return c / 2, true
	}
	return c, false
}

func Shrink[T any](s []T) []T {
	c, flag := check(cap(s), len(s))
	if !flag {
		return s
	}
	next := make([]T, 0, c)
	next = append(next, s...)
	return next
}
