package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	ns, err := net.LookupHost("www.baidu.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "--%s\n", n)
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, inter := range interfaces {
		fmt.Fprintf(os.Stdout, "--%s:%s\n", inter.HardwareAddr, inter.Name)
	}
}
