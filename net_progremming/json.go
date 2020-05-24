package main

import (
	"encoding/json"
	"fmt"
)

/**

 */
type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float64
}
func main() {
	gobook := Book{
		Title:       "Go语言编程",
		Authors:    []string{"we","de"},
		Publisher:   "dewei",
		IsPublished: false,
		Price:       23.2,
	}
	b, err := json.Marshal(gobook)
	if err != nil {
		return
	}
	fmt.Println(b)

	var book Book
	err = json.Unmarshal(b, &book)
	fmt.Println(book)
}