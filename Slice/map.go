package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

func Map[A, B any](s []A, f func(A) B) []B {
	return Slice.Map[A, B](s, f)
}
