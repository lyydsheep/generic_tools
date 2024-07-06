package Slice

import (
	"errors"
)

// Delete 删除切片中的指定元素
// 参数：s 切片，idx 要删除的元素索引
// 返回值：res 删除指定元素后的切片，err 错误信息
// 说明：如果索引超出切片范围，返回错误信息
func Delete[T any](s []T, idx int) ([]T, error) {
	if idx >= len(s) || idx < 0 {
		return nil, errors.New("index is invalid")
	}
	//idx之后的元素整体往前挪动一位
	for i := idx; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	return s[:len(s)-1], nil
}
