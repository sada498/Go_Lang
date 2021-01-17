package main

import "fmt"

func main() {
	postcodes := map[string]string{
		"Melbourne": "3000",
		"Preston":   "3072",
		"kew":       "3084",
	}
	fmt.Println(postcodes)

	ages := map[string]int{}

	ages["sada"] = 27
	ages["siva"] = 28
	ages["ramu"] = 25
	fmt.Println(ages)

	//to remove data fro map
	delete(postcodes, "kew")
	fmt.Println(postcodes)

}
