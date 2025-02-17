package main

import (
	"fmt"
	"port-scanner/scanner"
)

func main() {
	var target string
	var startPort, endPort int

	fmt.Print("Digite o IP ou domínio para escanear: ")
	fmt.Scan(&target)

	fmt.Print("Digite a porta inicial: ")
	fmt.Scan(&startPort)

	fmt.Print("Digite a porta final: ")
	fmt.Scan(&endPort)

	scanner.ScanPorts(target, startPort, endPort)
}
