package generic_tools

func toPtr[T any](x T) *T {
	return &x
}
