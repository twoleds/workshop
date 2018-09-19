package main

import (
	"fmt"
	"time"
)

func printMsg(msg string) {

	fmt.Println("entering printMsg function")

	defer func() {
		fmt.Println("leaving printMsg function")
	}()

	for i := 1; i <= 4; i++ {
		fmt.Printf("%d: %s\n", i, msg)
		time.Sleep(1 * time.Second)
	}

}

func main() {

	printMsg("foo")
	time.Sleep(5 * time.Second)

}
