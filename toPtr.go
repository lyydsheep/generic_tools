package generic_tools

func ToPtr[T any](x T) *T {
	return &x
}
