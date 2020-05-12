package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(0)
	} else {
		fmt.Println("The address is", addr.String())
		text, _ := addr.MarshalText()
		fmt.Println("address bytes", text)
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is", addr.String(),
		"\nDefault mask length is", bits,
		"\nLeading ones count is", ones,
		"\nMask is (hex)", mask.String(),
		"\nNetwork is", network.String())

	os.Exit(0)
}
