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
StringToInt("$5", "USD") // output = 500

DisplayFull(10, "USD") // output = "$0.10"

DisplayWithAlpha(0, "USD") // output = "USD 0.00"

DisplayNoSymbol(10, "USD") // output = "0.10"

DisplayWithDecimal(10, "USD") // output = "0.10"

TopCurrencies() // output = []string{"USD", "EUR", "GBP", "INR", "CRC", "VND", "HUF", "ILS", "CNY", "KRW", "NGN", "PYG", "PHP", "PLN", "THB", "UAH", "JPY"}

ListCurrencies([]string{"USD"}) // output = []Currency{{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}}
```

## Additional Functions
Here are additional functions that may come in handy to your currency needs.

```go
GetISOFromAlpha("USD") // output = Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "$", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}

GetISOCodeFromNumeric("840") // output = "840"

GetAlphaFromISONumeric("840") // output = "USD"

ConvertToStringWithDecimal(100, 2) // output = "1.00"

InsertDelimiter("0001", 3, ",") // output = "000,1"

SwapSymbolWithAlpha("$0.00", "$", "USD") // output = "USD 0.00"

IsNegative(-1) // output = true

FormatCurrency(1, Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}) // output = "$0.01"

FloatToInt(9.99, 2) // output = 999

IntToFloat(999, 2) // output = 9.99

PercentageFromInt(898, 56.7, 3, Round) // output = 509.166

PercentageFromFloat(11.11, 13, 4, Round) // output = 1.4443

MaskCard("4111111111111111") // output = "411111", "1111", "411111******1111"

MaskACHAccount("8114460248") // output = "81******48"

ValidLuhn("4111111111111111") // output = true

GetCardType("4111111111111111") // output = "visa"
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
