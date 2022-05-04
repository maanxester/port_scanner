package main

import (
	"fmt"
	"github.com/maanxester/port_scanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")
	//open := port.ScanPort("tcp", "localhost", 5555)
	//fmt.Println("Port Open: %t\n", open)

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
