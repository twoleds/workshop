package main

import "fmt"

func printMsg(msg string) {

	if msg == "" {
		panic("Empty input parameter msg")
	}

	fmt.Println(msg)

}

func main() {

	printMsg("foo")
	printMsg("")
	printMsg("bar")

}
