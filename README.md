# Dough
![Dough](/dough.svg)

"Roll Out" your golang currency issues with Dough.

[![Build Status](https://travis-ci.com/fluidpay/dough.svg?branch=master)](https://travis-ci.com/fluidpay/dough)

## Installation
```sh
$ go get -u github.com/fluidpay/dough
```

## Currency Functions
Here are the available functions to help deal with currency issues you may run into.

```go
StringToInt(num string, alpha string) (int, error)
StringToInt("$5", "USD") // output = 500

DisplayFull(num int, alpha string) (string, error)
DisplayFull(10, "USD") // output = "$0.10"

DisplayWithAlpha(num int, alpha string) (string, error)
DisplayWithAlpha(0, "USD") // output = "USD 0.00"

DisplayNoSymbol(num int, alpha string) (string, error)
DisplayNoSymbol(10, "USD") // output = "0.10"

DisplayWithDecimal(num int, alpha string) (string, error)
DisplayWithDecimal(10, "USD") // output = "0.10"

TopCurrencies() ([]Currency, error)
TopCurrencies() // output = []string{"USD", "EUR", "GBP", "INR", "CRC", "VND", "HUF", "ILS", "CNY", "KRW", "NGN", "PYG", "PHP", "PLN", "THB", "UAH", "JPY"}

ListCurrencies(list []string) ([]Currency, error)
ListCurrencies([]string{"USD"}) // output = []Currency{{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}}
```

## Additional Functions
Here are additional functions that may come in handy to your currency needs.

```go
GetISOFromAlpha(alpha string) (Currency, error)
GetISOFromAlpha("USD") // output = Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "$", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}

GetISOCodeFromNumeric(num string) (string, error)
GetISOCodeFromNumeric("840") // output = "840"

GetAlphaFromISONumeric(num string) (string, error)
GetAlphaFromISONumeric("840") // output = "USD"

ConvertToStringWithDecimal(num int, fraction int) string
ConvertToStringWithDecimal(100, 2) // output = "1.00"

ReverseString(str string) string
ReverseString("001") // output = "100"

InsertDelimiter(str string, group int, del string) string
InsertDelimiter("0001", 3, ",") // output = "000,1"

SwapSymbolWithAlpha(str string, sym string, alpha string) string
SwapSymbolWithAlpha("$0.00", "$", "USD") // output = "USD 0.00"

RemoveSymbol(str string, sym string) string
RemoveSymbol("$0.00", "$") // output = "0.00"

RemoveDelimiter(str string, del string) string
RemoveDelimiter("$0.00", ".") // output = "$000"

RemoveDecimal(str string, dec string) string
RemoveDecimal("$0.00", ".", "$000")

IsNegative(num int) bool
IsNegative(-1) // output = true

FormatCurrency(num int, ISO Currency) string
FormatCurrency(1, Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}) // output = "$0.01"

FloatToInt(amt float64, fraction int) int
FloatToInt(9.99, 2) // output = 999

IntToFloat(amt int, fraction int) float64
IntToFloat(999, 2) // output = 9.99

PercentageFromInt(amt int, percentage float64, fraction int) float64
PercentageFromInt(898, 56.7, 3) // output = 509.166

PercentageFromFloat(amt float64, percentage float64, fraction int) float64
PercentageFromFloat(11.11, 13, 4) // output = 1.4443
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)