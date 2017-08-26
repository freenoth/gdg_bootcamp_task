# Currency converter

[![Go Report Card](https://goreportcard.com/badge/github.com/freenoth/gdg_bootcamp_task)](https://goreportcard.com/report/github.com/freenoth/gdg_bootcamp_task)

Go CLI application for currencies converting.

A task for Golang Bootcamp Perm, August 26, 2017
info: https://www.meetup.com/GDG-Perm/events/242030053/

## Requirements

`go get "golang.org/x/net/html/charset"`

## Installing

`go get github.com/freenoth/gdg_bootcamp_task`

## Help info

```
This app provides converting  value in one currency to all others on some date.

Headers:
    Show info where currencies are got and courses date.

    EXAMPLE:
    GET     http://www.cbr.ru/scripts/XML_daily.asp
    Origin: Foreign Currency Market
    Date:   24.08.2017

Currencies table:
    The result of converting passed value to others currencies.

    EXAMPLE:
    #    CODE   NAME                                         VALUE
    0    AUD    Австралийский доллар                         1.1
    1    AZN    Азербайджанский манат                        1.5411341
    2    GBP    Фунт стерлингов Соединенного королевства     0.7140677

Specify value to convert:
    The default currency and value is 1 RUB.

    You can specify other values by args:
    '-c CURRENCY' to specify initial currency (uppercase currency code)
    '-v VALUE' to specify initial value (dotted float value)

    EXAMPLE:
    conveter -v 100 -c AUD

    #    CODE   NAME                                         VALUE
    0    AUD    Австралийский доллар                         100.00
    1    GBP    Фунт стерлингов Соединенного королевства     61.63
    ...

    '-d DATE' to specify date of currencies courses (DATE format is 'dd/mm/yyy')

    EXAMPLE:
    conveter -v 100 -c AUD -d 01/01/2007

    GET     http://www.cbr.ru/scripts/XML_daily.asp?date_req=01/01/2007
    Origin: Foreign Currency Market
    Date:   30.12.2006

    #    CODE   NAME                                         VALUE
    0    AUD    Австралийский доллар                         100.00
    1    GBP    Фунт стерлингов Соединенного королевства     40.31
```
