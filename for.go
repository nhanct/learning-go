package main

import (
	"fmt"
)

func main() {
	slice := []string{"a", "b", "c"}
	for i, v := range slice {
	fmt.Println(i,v)
	}

        mapX := map[string]string {
	"vn" : "Viet Nam",
	"en" : "English",
	}
	for x, y := range mapX {
	fmt.Println(x, y)
	}
	
}
