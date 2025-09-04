# Prompt-Powered Kickstart: Building a Beginner’s Toolkit for Go

## 1. Title & Objective
**Getting Started with Go – A Beginner’s Guide**

- **Chosen Technology:** Go (Golang)
- **Why Go?** Go is a statically typed, compiled language designed for simplicity and efficiency. It’s widely used for backend services, cloud infrastructure, and command-line tools.
- **End Goal:** Run a minimal "Hello World" Go program.

## 2. Quick Summary of Go
Go is an open-source programming language developed by Google. It is known for its simplicity, concurrency support, and fast compilation. Go is used in cloud services (e.g., Docker, Kubernetes), web servers, and CLI tools.

- **What is it?** A compiled, statically typed language for building scalable software.
- **Where is it used?** Backend development, cloud infrastructure, DevOps tools.
- **Real-world example:** Docker, Kubernetes, and Terraform are written in Go.

## 3. System Requirements
- **OS:** Linux, Mac, Windows
- **Tools/Editors:** VS Code (recommended), any text editor
- **Packages:** Go (golang) compiler

## 4. Installation & Setup Instructions
### Step 1: Install Go
#### On Linux (Debian/Ubuntu)
```bash
sudo apt update
sudo apt install golang-go
```
#### Verify Installation
```bash
go version
```
Expected output: `go version go1.x.x ...`

### Step 2: Clone or Create Project Directory
```bash
mkdir go-beginner-toolkit
cd go-beginner-toolkit
```

### Step 3: Create Hello World File
Create a file named `main.go` with the following content:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World from Go!")
}
```

## 5. Minimal Working Example
This program prints `Hello, World from Go!` to the terminal.

### Run the Program
```bash
go run main.go
```
Expected output:
```
Hello, World from Go!
```

## 6. AI Prompt Journal
- **Prompt used:**
  - "Give me a step-by-step guide to initialize a Go Hello World project on Linux."
- **AI’s response summary:**
  - Provided installation steps, project scaffolding, and code example.
- **Evaluation:**
  - Very helpful for quickly setting up Go and understanding the basics.

## 7. Common Issues & Fixes
- **Issue:** `go: command not found`
  - **Fix:** Ensure Go is installed and added to your PATH. Run `sudo apt install golang-go`.
- **Issue:** Permission denied when installing Go
  - **Fix:** Use `sudo` for installation commands.
- **Issue:** Typo in code (e.g., missing quotes or parentheses)
  - **Fix:** Double-check code syntax and rerun.

## 8. References
- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Moringa School AI Curriculum](https://ai.moringaschool.com/)
- [Stack Overflow Go Questions](https://stackoverflow.com/questions/tagged/go)

---

**README instructions and codebase included in this folder.**
