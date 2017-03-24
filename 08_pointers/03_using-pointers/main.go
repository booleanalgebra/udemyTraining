package main

import "fmt"

func main() {

	var memory_location uint64
	fmt.Println("Enter your value to be placed in memory")
	fmt.Scan(&memory_location) //putting the value in a memory location
	fmt.Println("Your value is", memory_location)
	fmt.Println("and your value is stored memory location", &memory_location)


	fmt.Println("Now, enter a different value to be stored in this memory location")
	var mem_pointer = &memory_location //the outcome will be the same if you use mem_pointer *uint64
	fmt.Scan(mem_pointer)
	fmt.Println("Your value is", *mem_pointer) //the *returns the value in the location here
	fmt.Println("and is stored in memory location", mem_pointer)


	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
