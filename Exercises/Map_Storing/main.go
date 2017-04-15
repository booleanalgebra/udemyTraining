package main

import "fmt"

func main() {


	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",

	}
	fmt.Println(elements)

	elements_2 := map[string]map[string]string{

		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name": "Lithium",
			"state":"solid",
		},

	}
	fmt.Println(elements_2["H"])
	fmt.Println(elements_2["He"])
	fmt.Println(elements_2["Li"])
}

