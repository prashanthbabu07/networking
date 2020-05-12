/*
	LookupPort
	The network argument is a string such as "tcp" or "udp", while the service is a string such as "telnet" or "domain" (for DNS).
	go run lookup_port.go tcp telnet
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr,
			"Usage: %s network-type service\n",
			os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port", port)
	os.Exit(0)
}
