package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

func Intersection[T comparable](s1 []T, s2 []T) (res []T) {
	return Slice.Intersection[T](s1, s2)
}
