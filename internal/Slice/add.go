package Slice

import "errors"

//Add 向切片中指定位置插入一个元素
//参数：s 切片，val 要插入的元素，idx 指定的位置
//返回值：插入元素后的切片以及错误信息
//注意：如果idx小于0，返回错误信息；如果idx大于等于切片长度，在切片末尾追加指定元素

func Add[T any](s []T, val T, idx int) ([]T, error) {
	if idx < 0 {
		return nil, errors.New("index is invalid")
	} else if idx > len(s) {
		idx = len(s)
	}
	var tmp T
	s = append(s, tmp)
	for i := len(s) - 1; i > idx; i-- {
		s[i] = s[i-1]
	}
	s[idx] = val
	return s, nil
}
