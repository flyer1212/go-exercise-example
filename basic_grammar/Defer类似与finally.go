package main

import (
	"fmt"
	"os"
)

/**

 */

func createFile(p string) *os.File{
	fmt.Println("createing")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("wiring")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}
func main() {
	f := createFile("C:\\Users\\liuZOZO\\Desktop\\goc\\defer.txt")
	defer closeFile(f) // 无论放到什么位置都会执行
	writeFile(f)

}