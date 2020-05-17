package main

/**
  不需要使用函数名
 */

func intSeq() func() int{
	i := 0
	return func() int {
		i +=1
		return i
	}
}
func main() {
	nextInt := intSeq()
	println(nextInt())
	println(nextInt())
	println(nextInt())

	newInts := intSeq()
	println(newInts())
}