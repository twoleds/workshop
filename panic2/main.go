package main

import "fmt"

func printMsg(msg string) {

	if msg == "" {
		panic("Empty input parameter msg")
	}

	fmt.Println(msg)

}

func printMsg2(msg string) {

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("Err: %s\n", p)
		}
	}()

	printMsg(msg)

}

func main() {

	printMsg2("foo")
	printMsg2("")
	printMsg2("bar")

}
