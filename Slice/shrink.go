package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

func Shrink[T any](s []T) []T {
	return Slice.Shrink[T](s)
}
