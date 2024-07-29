package main

import (
	"fmt"
	"github.com/lyydsheep/generic_tools/Slice"
)

func main() {

	// 测试Slice工具包
	fmt.Println(Slice.Map[int, int]([]int{1, 2, 3, 4}, func(i int) int {
		return i * 2
	}))
}
