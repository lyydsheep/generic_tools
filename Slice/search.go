package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

//Search 查找切片中的指定元素
//参数：s 切片， val 要查找的元素
//返回值：res 目标元素的下标的切片

func Search[T comparable](s []T, val T) (res []int) {
	return Slice.Search[T](s, val)
}
