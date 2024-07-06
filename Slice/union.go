package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

//Union 求两个切片的并集
//参数：s1第一个切片，s2第二个切片
//返回值：并集结果
//说明：结果不含重复元素

func Union[T comparable](s1, s2 []T) (res []T) {
	return Slice.Union[T](s1, s2)
}
