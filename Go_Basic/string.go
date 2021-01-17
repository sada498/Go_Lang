package main

import "fmt"

func main() {
	fmt.Println("string name sada")
	fmt.Println(`
	This is multi -line
	string, that can also contain "quotes".
	`)

	fmt.Println(":)")
	// unicode
	fmt.Println("\u2272")
}
