package main

import (
	"fmt"
)

func main() {
	var a int = 8 //create variable a with value 8, type int
	b := 9 // create variable b and init value 8
	b = 22 // assign varibable b that we'v init to value 8
	var color string = "yellow" // set variable color with string type and value yellow
	fmt.Println(b) //print b
	fmt.Println(a) //print a
	fmt.Println(color) //print color
	var mya [2]string  // create array mya with string type and lenght is 2
	mya[0] = "dcm" // index 0 is dcm
	mya[1] = "may" // index 1 is may
	fmt.Println(mya) // print dcm may
	myMap := make(map[string]int) // create map with key type string and value type int
	myMap["dcm"] = 10 // assing key dcm with value 10
	fmt.Println(myMap["dcm"]) // print them
	exMap := make(map[string]string) //
	exMap["focus"] = "for"
	exMap["prius"] = "toyota"
	fmt.Printf("%s", exMap["prius"])
	
	exMapCom := map[string]string {
	"focus" : "for",
	"prius" : "toyota",
	}
	fmt.Printf("%s", exMapCom["prius"])
}
