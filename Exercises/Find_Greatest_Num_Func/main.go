package main

import "fmt"

func main(){
	x := greatest_num(1, 4, 5, 9.1, 9.9, 9.99, 9.999, 10.001, 5)
	fmt.Println(x)
}
func greatest_num(z... float64) float64{
	var grt float64
	for _, value_from_z := range z {
		if value_from_z > grt {grt = value_from_z

			}
	}
	return grt
}