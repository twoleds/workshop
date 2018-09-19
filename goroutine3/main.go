package main

import (
	"fmt"
	"time"
)

func printMsg(msg string, sig chan<- struct{}) {

	for i := 1; i <= 4; i++ {
		fmt.Printf("%d: %s\n", i, msg)
		time.Sleep(1 * time.Second)
	}

	sig <- struct{}{}

}

func main() {

	sig := make(chan struct{})

	go printMsg("foo", sig)
	go printMsg("bar", sig)

	for i := 0; i < 2; i++ {
		<-sig
	}

}
