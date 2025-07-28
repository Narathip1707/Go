package main

import "fmt"

//This example shows declaring variables outside of a function, with the var keyword:

func main() {
	var (
		a int
		b int    = 7
		c string = "Hello"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
