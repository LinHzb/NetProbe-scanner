package scanner

import (
	"fmt"
	"net"
	"os"
)

// ValidateTarget checks if the provided address is a valid IP or domain
func ValidateTarget(target string) bool {
	if net.ParseIP(target) != nil {
		return true
	}
	if _, err := net.LookupHost(target); err == nil {
		return true
	}
	return false
}

// ConvertPortRange converts a range of ports into a list of ports
func ConvertPortRange(startPort, endPort int) ([]int, error) {
	if startPort > endPort {
		return nil, fmt.Errorf("start port cannot be greater than end port")
	}
	var ports []int
	for port := startPort; port <= endPort; port++ {
		ports = append(ports, port)
	}
	return ports, nil
}

// FormatResults formats and displays the scan results
func FormatResults(results map[int]bool) {
	for port, isOpen := range results {
		status := "closed"
		if isOpen {
			status = "open"
		}
		fmt.Printf("Port %d is %s\n", port, status)
	}
}

// SaveResults saves the scan results to a file
func SaveResults(results map[int]bool, filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	for port, isOpen := range results {
		status := "closed"
		if isOpen {
			status = "open"
		}
		_, err := file.WriteString(fmt.Sprintf("Port %d is %s\n", port, status))
		if err != nil {
			return err
		}
	}

	return nil
}
