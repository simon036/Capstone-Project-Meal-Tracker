package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount.")
		return
	}

	fmt.Print("From currency (e.g., USD): ")
	from, _ := reader.ReadString('\n')
	from = strings.ToUpper(strings.TrimSpace(from))

	fmt.Print("To currency (e.g., EUR): ")
	to, _ := reader.ReadString('\n')
	to = strings.ToUpper(strings.TrimSpace(to))

	
	apiKey := "5d16c1eba19d00628414c06f"
	apiURL := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/latest/%s", apiKey, from)
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching rates:", err)
		return
	}
	defer resp.Body.Close()

	var data struct {
		ConversionRates map[string]float64 `json:"conversion_rates"`
		Result          string             `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response.")
		return
	}
	if data.Result != "success" {
		fmt.Println("Invalid currency code or API error.")
		return
	}

	rate, ok := data.ConversionRates[to]
	if !ok {
		fmt.Println("Target currency not found.")
		return
	}

	converted := amount * rate
	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, converted, to)
}
