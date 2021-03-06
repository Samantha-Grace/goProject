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

//----------------------------------------

//this has no return value
// func addOne(x int) {
// 	x = x + 1
// }

// func main() {

// 	x := 5
// 	fmt.Println(x)

// 	var xPtr *int = &x
// 	fmt.Println(xPtr)

// 	addOne(x)
// 	fmt.Println(x)
// }

//--------------------------------------------------

//this has return value
// func addTwo (x int) int {
// }

//structs

// type position struct {
// 	x float32
// 	y float32
// }

// type badGuy struct {
// 	name   string
// 	health int
// 	pos    position
// }

// func whereIsBadGuy(b badGuy) {
// 	x := b.pos.x
// 	y := b.pos.y
// 	fmt.Println("(", x, ",", y, ")")
// }

// func main() {

// 	p := position{4, 2}

// 	fmt.Println(p.x)

// 	badguy := badGuy{"Susan the kitten", 100, p}
// 	whereIsBadGuy(badguy)
// }
