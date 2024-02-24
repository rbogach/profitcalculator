package main

import "fmt"

func main() {

	revenue := inputRevenue()
	expenses := inputExpenses()
	taxRate := inputTaxRate()

	ebt, profprofit, ratio := calculateFinancials(revenue, expenses, taxRate)

	println(ebt, profprofit, ratio)

}

func inputRevenue() float64 {
	var revenue float64
	fmt.Print("Please input revenue")
	fmt.Scan(&revenue)
	return revenue
}

func inputExpenses() float64 {
	var expenses float64
	fmt.Print("also ur expenses")
	fmt.Scan(&expenses)
	return expenses
}

func inputTaxRate() float64 {
	var taxRate float64
	fmt.Print("also ur TAXXX")
	fmt.Scan(&taxRate)
	return taxRate
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	profit := float64(earningBeforeTax) - taxRate
	ratio := float64(earningBeforeTax) / profit
	return earningBeforeTax, profit, ratio

}
