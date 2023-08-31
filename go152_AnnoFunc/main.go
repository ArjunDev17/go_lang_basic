package main

import (
	"fmt"
)

func main() {
	fmt.Println("Annonyms Function")
	//syntax
	/*
		func functionName(argument)return type{
			function body
			}
	*/
	func() {
		fmt.Println("Anonyms Function")
	}()
}
