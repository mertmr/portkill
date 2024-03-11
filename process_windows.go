// +build windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func killProcessOnPort(port int) error {
	// Find the process using the port.
	cmd := exec.Command("netstat", "-ano", "-p", "tcp")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error finding process for port %d: %w", port, err)
	}

	var pidStr string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, fmt.Sprintf(":%d", port)) {
			fields := strings.Fields(line)
			pidStr = fields[len(fields)-1]
			break
		}
	}

	if pidStr == "" {
		return fmt.Errorf("no process found listening on port %d", port)
	}

	// Convert the output to a PID.
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf("error converting process ID to number: %w", err)
	}

	// Kill the process.
	cmd = exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error killing process %d: %w", pid, err)
	}

	fmt.Printf("Process with PID %d on port %d has been killed.\n", pid, port)
	return nil
}