package Slice

import "github.com/lyydsheep/generic_tools/internal/Slices"

// Delete 删除切片中的指定元素
// 参数：s 切片，idx 要删除的元素索引
// 返回值：res 删除指定元素后的切片，err 错误信息
// 说明：如果索引超出切片范围，返回错误信息
func Delete[T any](s []T, idx int) (res []T, err error) {
	return Slices.Delete[T](s, idx)
}
