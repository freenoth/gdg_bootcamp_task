package main

import (
	"fmt"
	"gdg_bootcamp_task/args"
	"gdg_bootcamp_task/valutes"
)

func main() {
	curr, val, date := args.ParseFlags()

	fmt.Println("Welcome to Currency converter!\nprint --help for more information\n")
	vals := valutes.GetValuteCurrencies(date)

	valueInRoubles := vals.GetValueInRoubles(val, curr)

	fmt.Printf("%-5v%-7s%-45s%s\n", "#", "CODE", "NAME", "VALUE")
	for i, valute := range vals.Valute {
		fmt.Printf("%-5v%-7s%-45s%.2f\n", i, valute.CharCode, valute.Name, valute.FromRub(valueInRoubles))
	}
}
