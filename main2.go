package main

import "fmt"

//pointer
func addOne(num *int) {
	*num = *num + 1
}

func main() {

	x := 5
	fmt.Println(x)

	var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)
}
