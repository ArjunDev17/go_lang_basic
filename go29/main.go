package main

import "fmt"

func main() {
	fmt.Println("Channels")
	var firstChan chan int
	fmt.Println("Channel is :", firstChan)
	fmt.Printf("Channel types :%T", firstChan)
	fmt.Println("-------------------")
	ch := make(chan int)
	fmt.Println("Hello from main")
	go multiWithChan(ch)
	ch <- 10
	// if we declare ch value in before method  then we will get DEDLOCK Problem
	fmt.Println("-------------------------")
	ch3 := make(chan int)
	close(ch3)
	elem, ok := <-ch3
	fmt.Println("Hello God", elem, ok)
	//c:=make(chan string)
	ch4 := make(chan string)
	go initString(ch4)
	for {
		resp, ok := <-ch4
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", resp, ok)
	}

	fmt.Println("bye from amin")
	mych := make(chan string)
	mych <- "Ram"
	mych <- "Shita"
	mych <- "is"
	fmt.Println("***************************")
	fmt.Println("Mych data is :", mych)

	testch := make(chan string)
	go func() {
		testch <- "values1"
		testch <- "values2"
		testch <- "values3"
	}()
	for res := range testch {
		fmt.Println(res)
	}
}
func multiWithChan(ch chan int) {
	fmt.Println(2 * <-ch)
}
func initString(chanl chan string) {
	for v := 0; v < 3; v++ {
		chanl <- "Jai Shree Ram"
	}
}
