package Slices

//Filter 过滤切片中的指定元素
//参数：s 切片，val 要过滤的元素值
//返回值：res 过滤指定元素后的切片

func FilterVal[T comparable](s []T, val T) (res []T) {
	length := 0
	for i := 0; i < len(s); i++ {
		if s[i] != val {
			length++
		}
	}
	if length > 256 {
		res = make([]T, length, length+(length>>2))
	} else {
		res = make([]T, length, length+(length>>1))
	}
	for i, j := 0, 0; i < len(s); i++ {
		if s[i] != val {
			res[j] = s[i]
			j++
		}
	}
	return
}

// Filter 过滤切片中重复的元素
// 参数：s 切片
// 返回值：res 去重后的切片
func Filter[T comparable](s []T) (res []T) {
	m := make(map[T]bool, len(s))
	for _, v := range s {
		m[v] = true
	}
	res = make([]T, len(m))
	i := 0
	for k := range m {
		res[i] = k
		i++
	}
	return
}