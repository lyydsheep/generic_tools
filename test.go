package main

import (
	"fmt"
	"github.com/lyydsheep/generic_tools/Slice"
)

func main() {

	// 测试Slice工具包
	tmp := []int{1, 2, 3, 4, 5, 2, 2}
	res := Slice.Filter(tmp, 2)
	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}
