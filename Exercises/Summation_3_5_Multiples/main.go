package main

import "fmt"

func main(){
	var z int
	for i := 0; i<1000; i++  {
		if i%3==0 {z += i

		} else if i%5==0 {z += i

		}

	}
	fmt.Println(z)
}
