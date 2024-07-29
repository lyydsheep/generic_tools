package Slice

func Map[A, B any](s []A, f func(A) B) []B {
	res := make([]B, len(s))
	for i := range s {
		res[i] = f(s[i])
	}
	return res
}
