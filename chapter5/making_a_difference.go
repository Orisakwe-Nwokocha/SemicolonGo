package chapter5

import (
	"fmt"
	"semicolonGo/utils"
)

func TakeGlobalWarmingQuiz() {
	input := utils.Input{}
	correctAnswers := 0

	fmt.Println("Welcome to the Global Warming Quiz!")
	questions := getQuizQuestions()

	for row := 0; row < len(questions); row++ {
		fmt.Printf("\nQuestion %d: %s\n", row+1, questions[row][0])
		for column := 1; column <= 4; column++ {
			fmt.Printf("\t%s\n", questions[row][column])
		}

		fmt.Println("\nEnter your answer: ")
		userAnswer := input.NextLine()
		if userAnswer == questions[row][5] {
			correctAnswers++
		}
	}

	fmt.Printf("\nNumber of correct answers: %d\n", correctAnswers)

	if correctAnswers == 5 {
		fmt.Println("Excellent!")
	} else if correctAnswers == 4 {
		fmt.Println("Very good!")
	} else {
		fmt.Println("Time to brush up on your knowledge of global warming. Here is a starter website: https://en.wikipedia.org/wiki/Climate_change")
	}
}

func CalculateFairTax() {
	totalExpenses := getTotalExpenses()
	fairTax := getFairTax(totalExpenses)

	fmt.Printf("Total expenses: $%.2f\n", totalExpenses)
	fmt.Printf("Estimated FairTax: $%.2f\n", fairTax)
}

func getFairTax(totalExpenses float64) float64 {
	fairTaxRate := 0.30
	return totalExpenses * fairTaxRate
}

func getTotalExpenses() float64 {
	expenses := getExpenses()
	var totalExpenses float64

	for _, expense := range expenses {
		totalExpenses += expense
	}

	return totalExpenses
}

func getExpenses() map[string]float64 {
	input := utils.Input{}
	expenses := map[string]float64{}
	expenseCategories := []string{"housing", "food", "clothing", "transportation",
		"education", "health care", "vacations"}

	for _, category := range expenseCategories {
		fmt.Printf("Enter expenses for %s:\n$", category)
		expense, _ := input.NextFloat()
		expenses[category] = expense
	}

	return expenses
}
