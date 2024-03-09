package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const resultsFile = "results.txt"

func writeResultToFile(ebt, profprofit, ratio float64) {
	data := fmt.Sprintf("%sebt\n%d profit\n%s ratio\n", ebt, profprofit, ratio)

	os.WriteFile(resultsFile, []byte(data), 0644)

}

func getResults() (float64, error) {
	data, err := os.ReadFile(resultsFile)
	if err != nil {
		return 1000, errors.New("failed to find results file")

	}
	resultsText := string(data)
	results, err := strconv.ParseFloat(resultsText, 64)
	if err != nil {
		return 1000, errors.New("failed to parce stored results value")
	}
	return results, nil
}

func main() {

	revenue := inputRevenue()
	expenses := inputExpenses()
	taxRate := inputTaxRate()

	ebt, profprofit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT"+"%.1f\n", ebt)
	fmt.Printf("Porfit before Tax"+"%.1f\n", profprofit)
	fmt.Printf("Ratio"+"%.1f\n", ratio)
	writeResultToFile(ebt, profprofit, ratio)
	getResults()

}

func inputRevenue() float64 {
	var revenue float64
	for {
		fmt.Print("Please input revenue")
		fmt.Scan(&revenue)
		if revenue > 0 {
			//fmt.Println("Invalid amount")
			break
		}
		fmt.Println("Invalid amount. Revenue must be greater than 0.")
	}

	return revenue
}

func inputExpenses() float64 {
	var expenses float64
	for {
		fmt.Print("also ur expenses")
		fmt.Scan(&expenses)

		if expenses > 0 {
			///fmt.Println("Invalid expenses")
			break

		}
		fmt.Println("Invalid amount. Revenue must be greater than 0.")

	}
	return expenses
}

func inputTaxRate() float64 {
	var taxRate float64
	for {
		fmt.Print("also ur TAXXX")
		fmt.Scan(&taxRate)
		if taxRate > 0 {
			//fmt.Println("Invalid Tax rate")
			break
		}
		fmt.Println("Invalid amount. Revenue must be greater than 0.")
	}

	return taxRate
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	profit := float64(earningBeforeTax) - taxRate
	ratio := float64(earningBeforeTax) / profit
	return earningBeforeTax, profit, ratio

}
