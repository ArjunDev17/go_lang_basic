package main

import "fmt"

type Address struct {
	VillName string
	pinCode  int
}
type Employee struct {
	Name        string
	CompanyName string
	Age         int
	address     Address
}

func main() {
	fmt.Println("Nexted Struct")
	employee := Employee{
		Name:        "Arjun",
		CompanyName: "Falca",
		Age:         24,
		address: Address{
			VillName: "BahadurNagar",
			pinCode:  229306,
		},
	}
	fmt.Println("Employee data :", employee)
	fmt.Println("Anonymas Structure ")
	test := struct {
		fName string
		lName string
	}{
		fName: "Ram",
		lName: "Ji",
	}
	fmt.Println("God Details :", test)
}
