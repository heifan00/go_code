package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

/*
All 返回一个迭代器，在 Go 中它是一个具有特殊签名的函数。
*/
func (lst List[T]) All() iter.Seq[T] {
	// 迭代器函数接受另一个函数作为参数，按照约定称为yield（但名称可以是任意的）。
	// 它将为我们想要迭代的每个元素调用yield，并记录yield的返回值以防止可能的提前终止。
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

/*
迭代不需要底层数据结构，甚至不必是有限的！这是一个返回斐波那契数迭代器的函数：
只要yield 保持返回 true，它就会一直运行。
*/
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			fmt.Println("a:==", a)
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	for e := range lst.All() {
		fmt.Println(e)
	}

	all := slices.Collect(lst.All())
	fmt.Println("all:", all)
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
