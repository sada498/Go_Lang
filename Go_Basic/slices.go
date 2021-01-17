package main

import "fmt"

func main() {

	names := []string{}
	names = append(names, "sada")
	names = append(names, "ramu", "siva", "babu")
	fmt.Println(names)
	// with amke
	names2 := make([]string, 4)
	names2[0] = "sada"
	names2[1] = "ramu"
	names2[2] = "siva"
	names2[3] = "babu"
	// un comment to check the error index out of
	//names2[4] = "rao"
	fmt.Println(names2)
}
