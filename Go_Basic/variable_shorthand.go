package main

import "fmt"

func main() {
	// we can decler in diff way also
	//var myInt int = 16
	var name string
	myInt := 16
	// another way to decler a var
	name = "keith"

	var val, ok = "oh yes!", true
	fmt.Println("myInt:", myInt)
	fmt.Println("myInt times two:", myInt*2)
	fmt.Println("val is :", val)
	fmt.Println("ok is:", ok)
	fmt.Println("name is :", name)

}
