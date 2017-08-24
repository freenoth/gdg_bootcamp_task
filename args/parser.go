package args

import (
    "flag"
    "os"
    "fmt"
    "gdg_bootcamp_task/args/help"
    "time"
    "strconv"
)

func ParseFlags() (currency string, value float32, date *time.Time) {
    helpFlag := flag.Bool("help", false, "Show help")
    dateFlag := flag.String("d", "", "Specify a date for get currencies. Example: '31/08/2017'. (default now)")
    currencyFlag := flag.String("c", "RUB", "Specify a currency to convert. Examples: 'USD', 'EUR', 'RUB'.")
    valueFlag := flag.String("v", "1", "Specify a value to convert. Examples: '99.10', '100', '0.15'.")
    flag.Parse()

    if *helpFlag {
        help.ShowHelp()
        os.Exit(0)
    }

    currency = *currencyFlag

    val, err := strconv.ParseFloat(*valueFlag, 32)
    if err != nil {
        fmt.Println(err)
        flag.PrintDefaults()
        os.Exit(1)
    }
    value = float32(val)

    date = nil
    if len(*dateFlag) > 0 {
        d, err := time.Parse("02/01/2006", *dateFlag)
        if err != nil {
            fmt.Println(err)
            flag.PrintDefaults()
            os.Exit(1)
        }
        date = &d
    }

    return currency, value, date
}
