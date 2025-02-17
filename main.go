package main

import (
	"flag"
	"fmt"

	"github.com/LinHzb/go-port-scanner/cmd/cli"
	"github.com/LinHzb/go-port-scanner/scanner"
	"github.com/fatih/color"
)

func main() {
	reportFile := flag.String("report", "output/report.txt", "Report file path")
	flag.Parse()

	cli.DisplayASCIIArt("assets/ascii.txt")
	cli.DisplayMenu()

	choice := cli.GetUserChoice()
	switch choice {
	case "1":
		var target string
		var startPort, endPort int

		color.Cyan("🔍 Enter the IP or domain to scan: ")
		fmt.Scan(&target)

		if !scanner.ValidateTarget(target) {
			color.Red("❌ Invalid address.")
			return
		}

		color.Cyan("🔍 Enter the start port: ")
		fmt.Scan(&startPort)

		color.Cyan("🔍 Enter the end port: ")
		fmt.Scan(&endPort)

		ports, err := scanner.ConvertPortRange(startPort, endPort)
		if err != nil {
			color.Red("❌ %v", err)
			return
		}

		results := scanner.ConcurrentScan(target, ports, "tcp")
		scanner.FormatResults(results)

		err = scanner.SaveResults(results, *reportFile)
		if err != nil {
			color.Red("❌ Error saving report: %v", err)
		} else {
			color.Green("✅ Report saved successfully.")
		}
	case "2":
		color.Yellow("👋 Exiting...")
		return
	default:
		color.Red("❌ Invalid choice. Exiting...")
		return
	}
}
