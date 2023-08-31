package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array in depth")
	arr1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("Array data is :", arr1)
	arr := arr1
	arr[0] = 11
	fmt.Println("Array data is :", arr1)
	arr2 := &arr1
	arr2[0] = 99
	//arr1[0] = 12
	fmt.Println("Array data is :", arr1)
	displayAr(arr)

}
func displayAr(a [5]int) {

	for i := 0; i < 5; i++ {
		fmt.Println("Array dta is :", a[i])
	}
	for _, value := range a {
		fmt.Println("For each loop Array data is ", value)
	}
}
