package dough

import "errors"

// ErrorInvalidISO : returns an error for an invalid ISO code
var ErrorInvalidISO = errors.New("Invalid ISO Code")

// ErrorInvalidStringFormat : returns an error if trying to convert an invalid string value
var ErrorInvalidStringFormat = errors.New("Invalid String Format")

// ErrorInvalidISOFractionMatch : returns an error if fraction does not match ISO fraction
var ErrorInvalidISOFractionMatch = errors.New("Invalid ISO Fraction Match")

// ErrorUnableToFormatCurrency : returns an error for invalid currency formatting
var ErrorUnableToFormatCurrency = errors.New("Unable To Format Currency")

// ErrorUnableToFormatCurrencyFromString : returns an error for invalid formatting from a string
var ErrorUnableToFormatCurrencyFromString = errors.New("Unable To Format Currency From String")
