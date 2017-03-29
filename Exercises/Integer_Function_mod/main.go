package main

import "fmt"

func half(n, y int)(int, bool, bool) {
	fmt.Println("Enter a value")
	fmt.Scan(&n, &y) //added data entry
	return n / y, n % y == 0, n/y==1 //experimenting with 3 returns
}
func main(){
	h, even, equal := half(1,1) //why?
	fmt.Println(h, even, equal)
}
