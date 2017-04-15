package main

import "fmt"

func main(){
	fmt.Println("Enter Value to Factorial")
	var fv int
	fmt.Scan(&fv)
	f:= factorial1(fv)
	for n := range f {
		fmt.Println("Your Value is:", n)
	}
}

func factorial1(n int) chan int {
	out := make (chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}