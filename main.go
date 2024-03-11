package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No port number provided.")
		fmt.Println("Usage: portkill.exe <port_number>")
		os.Exit(1)
	}
	
	portStr := os.Args[1]
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Invalid port number.")
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err == nil {
		conn.Close()
		err = killProcessOnPort(port)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("No process is listening on port %d\n", port)
	}
}