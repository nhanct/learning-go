package main

import (
	"fmt"
)
func changeInt32 (x int32) {
    x = 10
    fmt.Println(x)
}


func changeStr (o *string) {
    *o = "hello world"
    fmt.Println(*o)
}
func main() {
	var x int32 = 22
	changeInt32(x)
	fmt.Println(x)
	y := 100
	z := &y
	fmt.Println(z)
	
	q  := "Howdy Partner"
	
	fmt.Println(q)
	changeStr(&q)
	fmt.Println(q)
}
