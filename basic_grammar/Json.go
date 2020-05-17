package main

import (
	"encoding/json"
	"fmt"
)

/**

 */

type Response1 struct {
	Page int
	Fruits[] string
}

type Response2 struct {
	Page int `json:"page"`
	Fruits []string `json:"fruits"`
}
func main() {
	res1 := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach","pear"}}
	res1Json, _ := json.Marshal(res1)
	fmt.Println(string(res1Json))

	str := `{"page":1, "fruits" :["apple", "peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}