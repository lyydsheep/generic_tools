package Slice

func Intersection[T comparable](s1 []T, s2 []T) (res []T) {
	s1 = Filter[T](s1)
	s2 = Filter[T](s2)
	m := make(map[T]bool, min(len(s1), len(s2)))
	for _, t := range s1 {
		m[t] = true
	}
	//res = make([]T, min(len(s1), len(s2)))
	i, length := 0, 0
	for _, t := range s2 {
		if _, ok := m[t]; ok {
			length++
		}
	}
	res = make([]T, length)
	for _, t := range s2 {
		if _, ok := m[t]; ok {
			res[i] = t
			i++
		}
	}
	return
}
