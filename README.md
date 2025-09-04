

# Go Currency Converter CLI – Beginner’s Toolkit

## 1. Title & Objective
**Title:** Go Currency Converter CLI  
**Objective:** Build a simple CLI tool in Go that converts currencies using real-time exchange rates from a public API. Learn Go basics, API calls, and environment variable management.

## 2. What You’ll Learn
- Go syntax and project structure
- CLI input handling
- HTTP requests and JSON parsing
- Using environment variables for API keys
- Error handling in Go

## 3. Prerequisites
- Go installed ([Download Go](https://go.dev/dl/))
- API key from [exchangerate-api.com](https://www.exchangerate-api.com/)
- Text editor (VS Code recommended)

## 4. Installation & Setup

Clone the repo
git clone <your-repo-url>
cd capstone

# Install dependencies
go mod tidy


## 5. Running the Project
go run main.go

Follow the prompts:
- Enter amount (e.g., 100)
- Enter source currency (e.g., USD)
- Enter target currency (e.g., EUR)

**Sample Output:**
```
100 USD = 92.50 EUR
```

## 6. Project Structure
capstone/
├── currency_converter.go  # Currency converter CLI
├── go_basics_demo.go      # Go basics demo (functions, loops, Hello World)
├── README.md              # This file
├── TOOLKIT.md             # Extended documentation
├── .env                   # API key (not included in repo)
├── go.mod, go.sum         # Go module files

## 7. AI Prompt Journal
- “How do I make a CLI app in Go that fetches data from an API?”
- “How do I use environment variables in Go?”
- “How do I parse JSON in Go?”

**AI Response:**  
Helped scaffold the CLI, handle errors, and securely use API keys.

## 8. Common Issues & Solutions

- **API key not found:**  
  Make sure `.env` exists and contains your key.
- **Invalid amount:**  
  Enter a valid number when prompted.
- **Go not installed:**  
  Download and install Go from the official site.

## 9. References

- [Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Exchangerate API Docs](https://www.exchangerate-api.com/docs)
- [Moringa AI Curriculum](https://ai.moringaschool.com/)

## 10. License & Contributing

MIT License.
