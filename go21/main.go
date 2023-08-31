// package main

// import (
// 	"fmt"

// )
// type stu struct{
// 	Name string
// 	USN string
// 	result struct
// }
// type result struct{

//		Marks int
//	}
//
//	func main() {
//		fmt.Println("Jai Shree Ram")
//		u:=stu{
//			Name:"Ram",
//			USN:"12"
//			result{
//				Marks:98
//			}
//		}
//		fmt.Println("Student info :"u)
//	}
package main

import "fmt"

type result struct {
	Marks int
}

type stu struct {
	Name   string
	USN    string
	Result result
}

func main() {
	fmt.Println("Jai Shree Ram")

	u := stu{
		Name: "Ram",
		USN:  "12",
		Result: result{
			Marks: 98,
		},
	}
	fmt.Println("Student info:", u)
}
