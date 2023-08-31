package main

import "fmt"

func main() {
	fmt.Println("SLice")
	arr := [...]string{"i", "am", "Learner", "for", "dev"}
	fmt.Println("Array is :", arr)
	mySlice := arr[1:4]
	fmt.Println("Slice is :", mySlice)
	fmt.Println("Size of SLice :", len(mySlice))
	fmt.Println("Capacity of SLice :", cap(mySlice))
	mySlice = append(mySlice, "Ram ji")
	fmt.Println("------------------------------------------")
	fmt.Println("Slice is :", mySlice)
	fmt.Println("Size of SLice :", len(mySlice))
	fmt.Println("Capacity of SLice :", cap(mySlice))
	mySlice = append(mySlice, "Ram ji")
	fmt.Println("------------------------------------------")
	fmt.Println("Slice is :", mySlice)
	fmt.Println("Size of SLice :", len(mySlice))
	fmt.Println("Capacity of SLice :", cap(mySlice))
	fmt.Println("------------------------------------------")
	fmt.Println("Slice using make keyword")
	slic1 := make([]int, 3)
	fmt.Println("Slice is ", slic1)

}
