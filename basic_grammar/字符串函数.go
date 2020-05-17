package main

import "fmt"

/**

 */
import s "strings"
var p  = fmt.Println

func main() {

	p("Contains",s.Contains("test", "es" ))
	p("Split", s.Split("a-b-c-d-e", "-"))
	p("ToLower:" , s.ToLower("TEst"))
	p("ToUpper: ", s.ToUpper("teSe"))
	p("Len", len("hello"))
	p("char", ("hello"[1] -'a' ))
}