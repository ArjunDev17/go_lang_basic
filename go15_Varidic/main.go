package main

import "fmt"

func main() {
	fmt.Println("Varidic Function")
	VaridicFun(1, 2, 3, 4)
}

// it is like your var argument method
func VaridicFun(v ...int) {
	fmt.Println("data are :", v)

}
