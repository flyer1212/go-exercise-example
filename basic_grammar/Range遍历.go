package main

import "fmt"

/**

  range遍历数组，遍历map,获取index
*/
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:",i)
		}
	}
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k,v)
	}
	// 打印unicode吗
	for i, c := range "go" {
		fmt.Println(i,c)
	}
}
