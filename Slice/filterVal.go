package Slice

import "github.com/lyydsheep/generic_tools/internal/Slice"

//FilterVal 过滤切片中的指定元素
//参数：s 切片，val 要过滤的元素值
//返回值：res 过滤指定元素后的切片

func FilterVal[T comparable](s []T, val T) (res []T) {
	return Slice.FilterVal[T](s, val)
}
