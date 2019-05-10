package dough

import "errors"

// ErrorInvalidISO : returns an error for an invalid ISO code
var ErrorInvalidISO = errors.New("Invalid ISO Code")

// ErrorUnableToFormatCurrency : returns an error for invalid currency formatting
var ErrorUnableToFormatCurrency = errors.New("Unable To Format Currency")

// ErrorUnableToFormatCurrencyFromString : returns an error for invalid formatting from a string
var ErrorUnableToFormatCurrencyFromString = errors.New("Unable To Format Currency From String")
