package main

import "fmt"

func main() {
	fmt.Println("Jai Shree Ram")
	a, _ := twoValu()
	c := a
	fmt.Println("Addition of two number is :", c)
}
func twoValu() (int, int) {
	return 12, 4
}
