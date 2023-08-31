package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array")
	var arr1 [2]int
	arr2 := [2]string{}
	fmt.Println("Array data are :", arr1)
	fmt.Println("Arrays data are :", arr2)
	arr1[0] = 12
	arr1[1] = 23
	arr2[0] = "Ram"
	arr2[1] = "ji"
	fmt.Println("Array data are :", arr1)
	fmt.Println("Arrays data are :", arr2)

	for i := 0; i < len(arr1); i++ {
		fmt.Println("Array data are :", arr1[i], reflect.TypeOf(arr1[0]))
	}
}
