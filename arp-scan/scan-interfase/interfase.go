package main

import (
	"fmt"
	"net"
)

func main() {
	// Get a list of all interfaces.
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, v := range ifaces {
		fmt.Println(v.Name,":Index is",v.Index,",and HardwareAddr is",v.HardwareAddr)
	}
}
