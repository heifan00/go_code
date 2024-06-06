package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, "len:", len(s), "cap:", cap(s))

	s[0] = "l"
	s[1] = "f"
	s[2] = "f"
	fmt.Println("s:", s)
	fmt.Println("get:", s[0])
	fmt.Println("len:", len(s))

	// 切片 append 需要接收新值
	s = append(s, "nihao")
	s = append(s, "shuai", "a")
	fmt.Println("apd:", s)

	// copy操作
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 左闭右开
	l := s[2:5]
	fmt.Println("l:", l)

	// 能获取到最大的
	l = s[:5]
	fmt.Println("l:", l)

	l = s[2:]
	fmt.Println("l:", l)

	// 初始化切片时就赋值
	t := []string{"g", "k", "d"}
	fmt.Println("t:", t)

	// 切片函数
	t2 := []string{"g", "k", "d"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// 切片可以组成多维数据结构。与多维数组不同，内部切片的长度可能会有所不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
