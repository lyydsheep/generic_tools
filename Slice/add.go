package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

func Add[T any](s []T, val T, idx int) ([]T, error) {
	return Slice.Add(s, val, idx)
}
