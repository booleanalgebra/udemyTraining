package main

import"fmt"

func main(){
	var large_number int
	var small_number int
	var mem_large_num = &large_number
	var mem_small_num = &small_number
	var remainder int
	fmt.Println("Please enter a large number")
	fmt.Scan(&large_number)
	fmt.Println("Now please enter a smaller number")
	fmt.Scan(&small_number)

	remainder = *mem_large_num % *mem_small_num
	fmt.Println("The remainder of the larger number divided by the smaller number is", remainder)
}
