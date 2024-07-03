package Slices

import (
	"errors"
)

// Delete 删除切片中的指定元素
// 参数：s 切片，idx 要删除的元素索引
// 返回值：res 删除指定元素后的切片，err 错误信息
// 说明：如果索引超出切片范围，返回错误信息
func Delete[T any](s []T, idx int) (res []T, err error) {
	if idx >= len(s) || idx < 0 {
		return nil, errors.New("index is invalid")
	}
	length := len(s)
	if length > 256 {
		res = make([]T, length-1, length+(length>>2))
	} else {
		res = make([]T, length-1, length+(length>>1))
	}
	for i, j := 0, 0; i < length; i++ {
		if i != idx {
			res[j] = s[i]
			j++
		}
	}
	return
}
