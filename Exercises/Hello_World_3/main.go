package main

import "fmt"

func main(){
	var name_input string
	fmt.Println("Please enter your name")
	fmt.Scan(&name_input)
	var name_output = &name_input
	fmt.Println("hello", *name_output)
}