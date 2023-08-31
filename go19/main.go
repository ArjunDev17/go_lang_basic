// package main

// import "fmt"

// type student struct {
// 	Name string
// 	Usn  string
// }

//	func main() {
//		fmt.Println("Method")
//		res := student{
//			Name: "Ram",
//			Usn:  "CR",
//		}
//		fmt.Println("Student details are :", res)
//		fmt.Printf("Student name is %s :", res.Name)
//		fmt.Printf("\nStudent USN is %s :", res.Usn)
//	}
//
//	func (s1 student)(s2 student)   {
//		s1.Name="arjun"
//		return s1.Name
//	}
package main

import "fmt"

type student struct {
	Name string
	Usn  string
}

// Method to update student's name
func (s *student) updateName(newName string) {
	s.Name = newName
}

func main() {
	fmt.Println("Method")

	res := student{
		Name: "Ram",
		Usn:  "CR",
	}
	ptr := &res
	fmt.Println("By Pointer :", ptr.Name)
	fmt.Println("Student details are:", res)

	res.updateName("Arjun")

	fmt.Printf("Student name is %s\n", res.Name)
	fmt.Printf("Student USN is %s\n", res.Usn)
}
