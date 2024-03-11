// +build !windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func killProcessOnPort(port int) error {
	cmd := exec.Command("lsof", "-i", fmt.Sprintf("tcp:%d", port), "-t")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error finding process for port %d: %w", port, err)
	}

	pidStr := strings.TrimSpace(string(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return fmt.Errorf("error converting process ID to number: %w", err)
	}

	cmd = exec.Command("kill", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error killing process %d: %w", pid, err)
	}

	fmt.Printf("Process with PID %d on port %d has been killed.\n", pid, port)
	return nil
}