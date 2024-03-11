package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	portFlag := flag.Int("port", 0, "Port number to kill the process on")
	flag.Parse()

	if *portFlag == 0 {
		fmt.Println("Please specify a port using the -port flag.")
		os.Exit(1)
	}

	port := *portFlag

	// Check if the port is in use.
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err == nil {
		conn.Close()
		// Kill the process using the port.
		err = killProcessOnPort(port)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("No process is listening on port %d\n", port)
	}
}