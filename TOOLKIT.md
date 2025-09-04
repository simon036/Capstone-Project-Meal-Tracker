# Prompt-Powered Kickstart: Go Currency Converter CLI Toolkit

## 1. Title & Objective
**Getting Started with Go – Currency Converter CLI**

- **Chosen Technology:** Go (Golang)
- **Why Go?** Go is a statically typed, compiled language designed for simplicity and efficiency. It’s widely used for backend services, cloud infrastructure, and command-line tools.
- **End Goal:** Build and run a CLI tool that converts currencies using real-time exchange rates from exchangerate-api.com.

## 2. Quick Summary of the Technology
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

go version

Expected output: `go version go1.x.x ...`

### Step 2: Clone Project Directory

git clone <your-repo-url>
cd capstone

### Step 3: Install Dependencies

go mod tidy

## 5. Minimal Working Example
This CLI program converts an amount from one currency to another using real-time exchange rates from exchangerate-api.com. The API key is hardcoded for demonstration purposes.

### Run the Program

go run currency_converter.go

Follow the prompts:
- Enter amount (e.g., 100)
- Enter source currency (e.g., USD)
- Enter target currency (e.g., KES)



## 6. AI Prompt Journal
- **Prompts used:**
  - "How do I make a CLI app in Go that fetches data from an API?"
  - "How do I parse JSON in Go?"
  - "How do I use an API key in Go?"
- **AI’s response summary:**
  - Helped scaffold the CLI, handle errors, and securely use API keys.
- **Evaluation:**
  - Very helpful for quickly setting up Go, understanding HTTP requests, and error handling.

## 7. Common Issues & Fixes
- **Issue:** `Invalid currency code or API error.`
  - **Fix:** Make sure you enter valid ISO currency codes (e.g., USD, KES, EUR).
- **Issue:** `Target currency not found.`
  - **Fix:** Check if the currency is supported by exchangerate-api.com.
- **Issue:** `Invalid amount.`
  - **Fix:** Enter a valid number when prompted.
- **Issue:** Network errors
  - **Fix:** Ensure you have an active internet connection.

## 8. References
- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Exchangerate API Docs](https://www.exchangerate-api.com/docs)
- [Moringa School AI Curriculum](https://ai.moringaschool.com/)
- [Stack Overflow Go Questions](https://stackoverflow.com/questions/tagged/go)

---

**README instructions and codebase included in this folder.**
