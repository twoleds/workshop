package main

import (
	"fmt"
	"os"
)

func write() (err error) {

	f, err := os.OpenFile("./workshop", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return
	}

	defer func() {
		err = f.Close()
	}()

	if _, err = f.WriteString("Hello world!"); err != nil {
		return
	}

	return

}

func main() {
	if err := write(); err != nil {
		fmt.Println(err)
	}
}
