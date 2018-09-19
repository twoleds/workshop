package main

import (
	"fmt"
	"time"
)

func printMsg(in <-chan string, sig chan<- struct{}) {

	msg := <-in

	for i := 1; i <= 4; i++ {
		fmt.Printf("%d: %s\n", i, msg)
		time.Sleep(1 * time.Second)
	}

	sig <- struct{}{}

}

func main() {

	sig, in := make(chan struct{}), make(chan string)

	for i := 0; i < 2; i++ {
		go printMsg(in, sig)
	}

	in <- "foo"
	in <- "bar"

	for i := 0; i < 2; i++ {
		<-sig
	}

}
