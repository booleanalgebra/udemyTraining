package main

import "fmt"

func main(){
	z := bool_exp(true, false, false, true, false, false)
	fmt.Println(z)

}
func bool_exp(a, b, c, d, e, f bool) bool{
	var  bool_eval bool
	if (a&&b)||(c&&d)||!(e&&f){
		bool_eval = true

	}

	return bool_eval

}