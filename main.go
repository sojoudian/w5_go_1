package main

import "fmt"

func main() {
	Greeting()
	advanceGreeting("M")
	fmt.Println(add(1, 2))
	c := add(1, 2)
	d := 5
	e := c + d
	fmt.Println(e)
	q, r := divide(10, 3)
	// var q int
	// q = 10
	// var q int = 10
	fmt.Printf("Quotient: %d, REminder: %d\n", q, r)
	q, r = 20, 3
	fmt.Printf("Quotient: %d, REminder: %d\n", q, r)

}
