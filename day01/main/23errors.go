package main

import (
	"errors"
	"fmt"
)

// 按照惯例，错误通常是最后一个返回值并且是 error 类型，它是一个内建的接口。
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New 使用给定的错误信息构造一个基本的 error 值。
		return -1, errors.New("can't work with 42")

	}

	// 返回错误值为 nil 代表没有错误。
	return arg + 3, nil
}

// 你还可以通过实现 Error() 方法来自定义 error 类型。 这里使用自定义错误类型来表示上面例子中的参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// 通过类型断言检查 e 是否可以被转换成 *argError 类型
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
