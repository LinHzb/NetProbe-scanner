package scanner

import (
	"fmt"
	"net"
	"time"
)

// ScanPorts verifica as portas em um endereço IP.
func ScanPorts(target string, startPort, endPort int) {
	fmt.Printf("Iniciando varredura de %d a %d em %s\n", startPort, endPort, target)
	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", target, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err != nil {
			fmt.Printf("Porta %d fechada\n", port)
			continue
		}
		conn.Close()
		fmt.Printf("Porta %d aberta\n", port)
	}
}
