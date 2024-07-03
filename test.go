package main

import (
	"fmt"
	"github.com/lyydsheep/generic_tools/Slice"
)

func main() {

	// 测试Slice工具包
	fmt.Printf("%+v", Slice.Intersection[int]([]int{1, 2, 3, 4, 5}, []int{1, 2, 2, 2, 5, 6}))
}
