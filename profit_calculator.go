package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var earningBeforeTax float64
	var profit float64

	revenue = inputRevenue()
	expenses = inputExpenses()
	taxRate = inputTaxRate()
	earningBeforeTax = earningBeforeTaxF(revenue, expenses)
	profit = profitF(earningBeforeTax, taxRate)

	println(earningBeforeTax)
	println(profit)
	println(ratioF(earningBeforeTax, profit))

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

func earningBeforeTaxF(revenue float64, expenses float64) float64 {
	earningBeforeTax := revenue - expenses
	println("Earnings before Tax")
	return earningBeforeTax

}

func profitF(earningBeforeTax float64, taxRate float64) float64 {
	profit := float64(earningBeforeTax) - taxRate
	println("Profit -")
	return profit
}

func ratioF(earningBeforeTax float64, profit float64) float64 {
	ratio := float64(earningBeforeTax) / profit

	println("Ratio is")
	return ratio

}
