package main

import (
	"errors"
	"fmt"
	"os"
)

const resultsFile = "results.txt"

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return

	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("TaxRate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profprofit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT"+"%.1f\n", ebt)
	fmt.Printf("Porfit before Tax"+"%.1f\n", profprofit)
	fmt.Printf("Ratio"+"%.1f\n", ratio)
	storeResults(ebt, profprofit, ratio)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	profit := float64(earningBeforeTax) - taxRate
	ratio := float64(earningBeforeTax) / profit
	return earningBeforeTax, profit, ratio

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("It must be a positive number")
	}

	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT	: %.1f\nProfit: %.1f\nRatio %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)

}
