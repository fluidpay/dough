package currency

import "errors"

// errorInvalidISO : returns an error for an invalid ISO code
var errorInvalidISO = errors.New("Error: Invalid ISO Code")

// errorUnableToFormatCurrency : returns an error for invalid currency formatting
var errorUnableToFormatCurrency = errors.New("Error: Unable To Format Currency")

// errorUnableToFormatCurrencyFromString : returns an error for invalid formatting from a string
var errorUnableToFormatCurrencyFromString = errors.New("Error: Unable To Format Currency From String")
