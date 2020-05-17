package main

import "fmt"

/**
  必须使用make创建

  二维数组中 每个维度长度相同
  slice不需要
*/
func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get", s[2])
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	t := []int{32, 23}
	fmt.Println("dcl:", t)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("apd", s)

	fmt.Println("sl1", s[2:5])
	fmt.Println("sl2", s[:5])
	fmt.Println("sl3", s[2:])

	twoD := make([][]int, 3)
	for i := 1; i < 9; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 1; j < innerLen; j++ {
			twoD[i][j] = i +j
		}
	}
	fmt.Println("2d", twoD)
}
