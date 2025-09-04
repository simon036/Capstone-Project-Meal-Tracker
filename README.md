# Go Meal Tracker CLI – Beginner’s Toolkit

## 1. Title & Objective
**Title:** Go Meal Tracker CLI  
**Objective:** Build a simple CLI tool in Go that helps users log meals, calories, and meal times, and view daily summaries. Learn Go basics, CLI input, and struct usage.

## 2. What You’ll Learn
- Go syntax and project structure
- CLI input handling
- Structs and slices
- Basic error handling in Go

## 3. Prerequisites
- Go installed ([Download Go](https://go.dev/dl/))
- Text editor (VS Code recommended)

## 4. Installation & Setup

# Clone the repo
git clone <your-repo-url>
cd capstone

# Install dependencies
go mod tidy


## 5. Running the Project

go run meal_tracker.go

Follow the prompts:
- Add a meal (name, calories, time)
- List all meals
- Show total calories for the day
- Exit

**Sample Output:**

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

## 6. Project Structure

capstone/
├── meal_tracker.go      # Meal tracker CLI
├── README.md            # This file
├── TOOLKIT.md           # Extended documentation
├── go.mod, go.sum       # Go module files
```

## 7. AI Prompt Journal

- “How do I make a CLI app in Go that stores and lists data?”
- “How do I use structs and slices in Go?”
- “How do I handle user input in Go?”

**AI Response:** 
Helped scaffold the CLI, handle errors, and use Go data structures.

## 8. Common Issues & Solutions

- **Invalid calories:**  
  Enter a valid number when prompted.
- **Go not installed:**  
  Download and install Go from the official site.

## 9. References

- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Moringa AI Curriculum](https://ai.moringaschool.com/)

## 10. License & Contributing

MIT License. 
