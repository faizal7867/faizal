package main

import "fmt"

func channels(value chan<- int) {
	for i := 0; i < 5; i++ {
		value <- i
	}
	close(value)
}

func channel(value <-chan int) {
	for value := range value {
		fmt.Println(value)
	}
}
func main() {
	ch := make(chan int)
	go channels(ch)
	go channel(ch)
}
