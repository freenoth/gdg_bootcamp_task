package valutes

import (
    "strings"
    "strconv"
    "fmt"
    "net/http"
    "encoding/xml"
    "golang.org/x/net/html/charset"
    "time"
    "os"
)

// ?date_req=23/08/2017
const valuteUrl = "http://www.cbr.ru/scripts/XML_daily.asp"

const RUB = "RUB"

type Valute struct {
    NumCode  int32
    CharCode string
    Nominal  int32
    Name     string
    Value    string
}

type ValCurs struct {
    XMLName xml.Name `xml:"ValCurs"`
    Date    string   `xml:"Date,attr"`
    Name    string   `xml:"name,attr"`
    Valute  []Valute
}

// Return value of current valute converted to roubles
func (v *Valute) ToRub(value float32) float32 {
    return value / float32(v.Nominal) * v.GetValue()
}

// Return value of current valute converted from roubles
func (v *Valute) FromRub(value float32) float32 {
    return value / v.GetValue() * float32(v.Nominal)
}

// Parse xml value with comma to initial float
func (v *Valute) GetValue() float32 {
    dotted_value := strings.Replace(v.Value, ",", ".", -1)
    val, _ := strconv.ParseFloat(dotted_value, 32)
    return float32(val)
}

func (vc *ValCurs) GetValueInRoubles(value float32, currency string) float32 {
    if currency == RUB {
        return value
    }

    for _, valute := range vc.Valute {
        if valute.CharCode == currency {
            return valute.ToRub(value)
        }
    }

    fmt.Println(fmt.Sprintf("no data for currency %s on %s", currency, vc.Date))
    os.Exit(1)
    return float32(0)
}

func GetValuteCurrencies(date *time.Time) *ValCurs {
    url := valuteUrl
    if date != nil {
        url = url + fmt.Sprintf("?date_req=%s", date.Format("02/01/2006"))
    }

    fmt.Printf("%-8s%s\n", "GET ", url)
    response, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer response.Body.Close()
    decoder := xml.NewDecoder(response.Body)
    decoder.CharsetReader = charset.NewReaderLabel

    var valCurs ValCurs
    err = decoder.Decode(&valCurs)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Printf("%-8s%s\n", "Origin:", valCurs.Name)
    fmt.Printf("%-8s%s\n", "Date:", valCurs.Date)
    fmt.Println("")

    valCurs.Valute = append(valCurs.Valute, Valute{
        NumCode:  1,
        CharCode: RUB,
        Nominal:  1,
        Name:     "Российский рубль",
        Value:    "1",
    })

    return &valCurs
}
