package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("JAi Shree Ram")
	MySLice := make([]int, 5)
	MySLice[0] = 99
	MySLice[1] = 88
	MySLice[2] = 77
	fmt.Println("Mydata is ", MySLice)
	MySLice = append(MySLice, 12, 4, 1, 66)
	sort.Ints(MySLice)
	fmt.Println("Sorted Data arre :", MySLice)
	fmt.Println("Sorted Data arre :", sort.IntsAreSorted(MySLice))

}
