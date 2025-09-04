package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Meal struct {
	Name     string
	Calories int
	Time     string
}

func main() {
	var meals []Meal
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMeal Tracker CLI")
		fmt.Println("1. Add meal")
		fmt.Println("2. List meals")
		fmt.Println("3. Show total calories")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			fmt.Print("Enter meal name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter calories: ")
			calStr, _ := reader.ReadString('\n')
			calStr = strings.TrimSpace(calStr)
			calories, err := strconv.Atoi(calStr)
			if err != nil {
				fmt.Println("Invalid calories.")
				continue
			}

			fmt.Print("Enter time (e.g., breakfast, lunch, dinner): ")
			time, _ := reader.ReadString('\n')
			time = strings.TrimSpace(time)

			meals = append(meals, Meal{Name: name, Calories: calories, Time: time})
			fmt.Println("Meal added!")

		case 2:
			fmt.Println("\nMeals:")
			for i, meal := range meals {
				fmt.Printf("%d. %s (%d cal) - %s\n", i+1, meal.Name, meal.Calories, meal.Time)
			}

		case 3:
			total := 0
			for _, meal := range meals {
				total += meal.Calories
			}
			fmt.Printf("Total calories: %d\n", total)

		case 4:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option.")
		}
	}
}
