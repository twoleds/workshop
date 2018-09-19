package main

import (
	"fmt"
	"time"
)

func printMsg(msg string) {
	for i := 1; i <= 4; i++ {
		fmt.Printf("%d: %s\n", i, msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printMsg("foo")
	go printMsg("bar")
	time.Sleep(5 * time.Second)
}
