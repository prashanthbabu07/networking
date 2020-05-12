/* LookupHost
 */

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	canonicalName, err := net.LookupCNAME(name)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(2)
	}
	fmt.Println("canonical name", canonicalName)

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(3)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
