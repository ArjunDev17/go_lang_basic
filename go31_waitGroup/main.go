package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d starting \n",id)
}
func main() {
	fmt.Println("Wait Group")
	
}
