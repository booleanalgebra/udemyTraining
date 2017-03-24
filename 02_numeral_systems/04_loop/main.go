package main

import "fmt"

//import "math"

func main() {
	for i := 0; i < 1000000; i++ {
		fmt.Printf("%d \t %b \t %#X \n", i+100, i, i)
	}
}
