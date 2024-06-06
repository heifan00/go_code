package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 8

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("k1:", v1)

	v3 := m["k3"]
	fmt.Println("k3:", v3)

	fmt.Println("len:", len(m))

	// 删除键值对
	delete(m, "k2")
	fmt.Println("del:", m)

	clear(m)
	fmt.Println("clear:", m)

	// 从映射中获取值时，可选的第二个返回值指示该键是否存在于映射中。这可以用于消除缺失键和零值（如0或“”）键之间的歧义。在这里，我们不需要值本身，所以我们使用空白标识符_忽略了它。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	aaa, prs := m["k2"]
	fmt.Println("prs:", prs, "aaa", aaa)

	// 初始化赋值
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")

	}
}
