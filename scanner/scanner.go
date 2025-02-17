package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPortTCP checks if a specific TCP port on the target is open
func ScanPortTCP(target string, port int, results chan<- int) {
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err != nil {
		results <- port
		return
	}
	conn.Close()
	results <- port
}

// ScanPortUDP checks if a specific UDP port on the target is open
func ScanPortUDP(target string, port int, results chan<- int) {
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("udp", address, 1*time.Second)
	if err != nil {
		results <- port
		return
	}
	conn.Close()
	results <- port
}

// ConcurrentScan scans multiple ports concurrently
func ConcurrentScan(target string, ports []int, protocol string) map[int]bool {
	results := make(map[int]bool)
	resultsChan := make(chan int)

	for _, port := range ports {
		if protocol == "tcp" {
			go ScanPortTCP(target, port, resultsChan)
		} else if protocol == "udp" {
			go ScanPortUDP(target, port, resultsChan)
		}
	}

	for range ports {
		port := <-resultsChan
		results[port] = true
	}

	return results
}
