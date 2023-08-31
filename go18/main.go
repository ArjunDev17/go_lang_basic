package main

import "fmt"

func msg(str string) {
	fmt.Println(str)
}
func project() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}
func main() {
	fmt.Println("CLouser function")
	msg("Jai Shree Ram")
	c := project()
	fmt.Println("Value is ", c())
	fmt.Println("Value is ", c())
	fmt.Println("Value is ", c())
}
