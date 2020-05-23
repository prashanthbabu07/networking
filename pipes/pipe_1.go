package main

import (
	"fmt"
	"os"
)

func main() {
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Size() > 0 {
		fmt.Println("there is something to read", fi.Size())
	} else {
		fmt.Println("stdin is empty")
	}

}
