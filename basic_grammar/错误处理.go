package main

import (
	"errors"
	"fmt"
)

/**
  异常信息用error类型，
  自定义异常 需要实现Error()方法

  自定义错误类型，需要定义类型断言，才能知道哪里的类型
*/

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with it"}
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
	for _, i := range [] int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}
	_, e := f2(42)
	// 如果 e 是 argError的类型
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
