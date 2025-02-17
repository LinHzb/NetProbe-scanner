### README.md

```markdown
# go-port-scanner

## Description

This project is a port scanner written in Go. It allows you to scan a range of ports on a specific IP or domain to check which ports are open or closed. The project also includes a command-line interface (CLI) with ASCII art and menu options.

## Features

- Display ASCII art at the start of the program.
- Interactive menu with options to start the scan or exit.
- User input for IP or domain and port range.
- Target validation (IP or domain).
- TCP port scanning.
- Formatting and displaying scan results.
- Saving scan results to a report file.
- Use of colors and emojis for a more friendly and visually appealing interface.

## Project Structure

```plaintext
port-scanner/
├── main.go
├── scanner/
│   ├── scanner.go
│   └── utils.go
├── cmd/
│   └── cli/
│       ├── menu.go
│       └── ascii.go
├── output/
│   └── report.txt
├── assets/
│   └── ascii.txt
└── go.mod
```

## How to Build and Run

### Prerequisites

- Go 1.22.2 or higher installed.

### Steps

1. Clone the repository:
   ```sh
   git clone https://github.com/LinHzb/go-port-scanner.git
   cd go-port-scanner
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Build the program:
   ```sh
   go build -o port-scanner
   ```

4. Run the program:
   ```sh
   ./port-scanner
   ```

## Usage

When you run the program, it will display the ASCII art and the menu. You can then choose an option:

```plaintext
███╗   ██╗███████╗████████╗██████╗ ██████╗  ██████╗ ██████╗ ███████╗
████╗  ██║██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██╔═══██╗██╔══██╗██╔════╝
██╔██╗ ██║█████╗     ██║   ██████╔╝██████╔╝██║   ██║██████╔╝█████╗  
██║╚██╗██║██╔══╝     ██║   ██╔═══╝ ██╔══██╗██║   ██║██╔══██╗██╔══╝  
██║ ╚████║███████╗   ██║   ██║     ██║  ██║╚██████╔╝██████╔╝███████╗
╚═╝  ╚═══╝╚══════╝   ╚═╝   ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝
Welcome to the Port Scanner
1. 🚀 Start Scan
2. ❌ Exit
Enter your choice: 1
🔍 Enter the IP or domain to scan: 8.8.8.8
🔍 Enter the start port: 80
🔍 Enter the end port: 85
```

The program will then start scanning the ports on the specified IP and port range and display the results with colors and emojis.

## Contribution

If you want to contribute to this project, follow the steps below:

1. Fork the repository.
2. Create a branch for your feature (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contact

For more information, contact [lingithz@gmail.com](mailto:lingithz@gmail.com).
```
