package main

import (
	"fmt"

	"github.com/freenoth/gdg_bootcamp_task/pkg/args"
	"github.com/freenoth/gdg_bootcamp_task/pkg/valutes"
)

func main() {
	curr, val, date := args.ParseFlags()

	fmt.Println("Welcome to Currency converter!\nprint --help for more information")
	fmt.Println("")
	vals := valutes.GetValuteCurrencies(date)

	valueInRoubles := vals.GetValueInRoubles(val, curr)

	fmt.Printf("%-5v%-7s%-45s%s\n", "#", "CODE", "NAME", "VALUE")
	for i, valute := range vals.Valute {
		fmt.Printf("%-5v%-7s%-45s%.2f\n", i, valute.CharCode, valute.Name, valute.FromRub(valueInRoubles))
	}
}
