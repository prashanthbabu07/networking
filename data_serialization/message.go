package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Message struct {
	Type uint
	Data string
}

func main() {
	fmt.Println([]byte("Hello"))

	message := &Message {
		Type: 1,
		Data: "hello message",
	}
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println(jsonBytes)
}
