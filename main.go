package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)
	//xPtr := &x

	var xPtr *int = &x
	fmt.Println(xPtr)

}

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
