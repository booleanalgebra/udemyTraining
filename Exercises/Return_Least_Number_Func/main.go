package main

import "fmt"

func main()  {
	d:=greatest_num( 3, 5, 9, 1, 2, 2, 1, 1.1, .01)//first find greatest
	x:= least_num(d, 3, 5, 9, 1, 2, 2, 1, 1.1, .01)//use same list
	fmt.Println(x)

}
func least_num(grt float64, y... float64) float64{
	var least = grt //pass greatest found value
	for _, value_from_y := range y {

		if value_from_y < least {least = value_from_y//check each value, if it's less than greatest, store it to least and keep checking

		}
	}
	return least
}
func greatest_num(z... float64) float64 {
	var grt float64
	for _, value_from_z := range z {
		if value_from_z > grt {
			grt = value_from_z

		}
	}
	return grt
}