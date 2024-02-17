package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64 = 30.0

	fmt.Print("Please input revenue")
	fmt.Scan(&revenue)
	fmt.Print("also ur expenses")
	fmt.Scan(&expenses)

	earningBeforeTax := revenue - expenses
	profit := float64(earningBeforeTax) - taxRate
	ratio := float64(earningBeforeTax) / profit
	println("Earnings before Tax")
	println(earningBeforeTax)
	println("Profit -")
	println(profit)
	println("Ratio is")
	println(ratio)

}
