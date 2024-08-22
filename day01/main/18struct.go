package main

import "fmt"

type person struct {
	name string
	age  int
}

// newPerson 使用给定的名字构造一个新的 person 结构体.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 24

	// 您可以安全地返回指向局部变量的指针， 因为局部变量将在函数的作用域中继续存在。
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Lee", age: 30})

	// 省略的字段将被初始化为零值。
	fmt.Println(person{name: "HiF"})

	// & 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 38})

	// 在构造函数中封装创建新的结构实例是一种习惯用法
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}
