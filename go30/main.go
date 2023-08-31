package main

import "fmt"

func main() {
	fmt.Println("It is Biderctional and")
	//creating a biderectional channel
	chan1 := make(chan string)
	chan2 := make(chan string)
	go sending(chan1)
	valueFromChannel := <-chan1
	fmt.Println("value from channel are :", valueFromChannel)
	go recieving(chan2)
	chan2 <- valueFromChannel
	go convert(chan1)
	fmt.Println(<-chan1)

}
func convert(s chan<- string) {
	s <- "Some value"
}
func sending(s chan string) {
	s <- "Go ram"
}
func recieving(s chan string) {
	fmt.Println(<-s)
}
