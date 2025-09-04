# Prompt-Powered Kickstart: Go Meal Tracker CLI Toolkit

## 1. Title & Objective
**Getting Started with Go – Meal Tracker CLI**

- **Chosen Technology:** Go (Golang)
- **Why Go?** Go is a statically typed, compiled language designed for simplicity and efficiency. It’s widely used for backend services, cloud infrastructure, and command-line tools.
- **End Goal:** Build and run a CLI tool that helps users log meals, calories, and meal times, and view daily summaries.

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
```bash
go version
```
Expected output: `go version go1.x.x ...`

### Step 2: Clone Project Directory
```bash
git clone <your-repo-url>
cd capstone
```

### Step 3: Install Dependencies
```bash
go mod tidy
```

## 5. Minimal Working Example
This CLI program lets users add meals, list them, and see total calories for the day—all locally, no API key or internet required.

### Run the Program
```bash
go run meal_tracker.go
```
Follow the prompts:
- Add a meal (name, calories, time)
- List all meals
- Show total calories for the day
- Exit

**Sample Output:**
```
Meal Tracker CLI
1. Add meal
2. List meals
3. Show total calories
4. Exit
Choose an option: 1
Enter meal name: Banana
Enter calories: 100
Enter time (e.g., breakfast, lunch, dinner): breakfast
Meal added!
```

## 6. AI Prompt Journal
- **Prompts used:**
  - "How do I make a CLI app in Go that stores and lists data?"
  - "How do I use structs and slices in Go?"
  - "How do I handle user input in Go?"
- **AI’s response summary:**
  - Helped scaffold the CLI, handle errors, and use Go data structures.
- **Evaluation:**
  - Very helpful for quickly setting up Go, understanding CLI input, and error handling.

## 7. Common Issues & Fixes
- **Issue:** `Invalid calories.`
  - **Fix:** Enter a valid number when prompted.
- **Issue:** Go not installed
  - **Fix:** Download and install Go from the official site.

## 8. References
- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Moringa School AI Curriculum](https://ai.moringaschool.com/)
- [Stack Overflow Go Questions](https://stackoverflow.com/questions/tagged/go)

---

**README instructions and codebase included in this folder.**
