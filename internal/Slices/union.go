package Slices

//Union 求两个切片的并集
//参数：s1第一个切片，s2第二个切片
//返回值：并集结果
//说明：结果不含重复元素

func Union[T comparable](s1, s2 []T) (res []T) {
	m := make(map[T]bool)
	for i := range s1 {
		m[s1[i]] = true
	}
	for i := range s2 {
		m[s2[i]] = true
	}
	res = make([]T, len(m))
	j := 0
	for k := range m {
		res[j] = k
		j++
	}
	return
}
