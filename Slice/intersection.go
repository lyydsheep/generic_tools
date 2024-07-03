package Slice

import "github.com/lyydsheep/generic_tools/internal/Slices"

func Intersection[T comparable](s1 []T, s2 []T) (res []T) {
	return Slices.Intersection[T](s1, s2)
}
