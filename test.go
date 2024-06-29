package main

import (
	"fmt"
	"github.com/lyydsheep/generic_tools/Slice"
)

func main() {

	// 测试Slice工具包
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{4, 5, 7, 6}
	//fmt.Printf("%+v", )
	fmt.Printf("%+v\n", Slice.Union[int](s1, s2))
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)
}
