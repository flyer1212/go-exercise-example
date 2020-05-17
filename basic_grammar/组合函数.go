package main

import (
	"fmt"
	"strings"
)

/**
  在一个函数里面，传递另一个函数
*/

func Index(vs [] string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func AnyContains(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs[] string, f func(string) bool) bool{
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
// 直接传唤为大小写
func Map(vs[] string, f func(string) string)[] string{
	vsm := make([]string, len(vs))
	for i, v:= range vs{
		vsm[i] = f(v)
	}
	return vsm
}
func main() {
	vs :=[]string{"peach", "apple","pearr"}
	fmt.Println(Any(vs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))
	fmt.Println(Map(vs, strings.ToUpper))
}
