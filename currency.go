package currency

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ConvertValueFromString : returns a uint from a string value
func ConvertValueFromString(num string) (uint, error) {
	reg := regexp.MustCompile("[0-9]+")
	str := reg.FindAllString(num, -1)
	newStr := strings.Join(str, "")
	output, err := strconv.ParseUint(newStr, 10, 64)
	if err != nil {
		return 0, errorUnableToFormatCurrencyFromString
	}
	return uint(output), nil
}

// ConvertToStringFull : returns currency string with full formatting
func ConvertToStringFull(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	output := formatCurrency(num, ISO)
	return output, nil
}

// ConvertToStringAlpha : returns currency string formatting with ISO alpha code instead of symbol
func ConvertToStringAlpha(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := formatCurrency(num, ISO)
	output := swapSymbolWithAlpha(currency, ISO.Symbol, ISO.Alpha)
	return output, nil
}

// ConvertToStringNoSymbol : returns currency string without currency symbol
func ConvertToStringNoSymbol(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := formatCurrency(num, ISO)
	output := removeSymbol(currency, ISO.Symbol)
	return output, nil
}

// ConvertToStringDecimal : returns currency string without currency symbol
func ConvertToStringDecimal(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := formatCurrency(num, ISO)
	currency = removeSymbol(currency, ISO.Symbol)
	output := removeDelimiter(currency, ISO.Delimiter)
	return output, nil
}

// ConvertToStringRaw : returns a raw value as a string
func ConvertToStringRaw(num uint, alpha string) (string, error) {
	_, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}

	output := fmt.Sprint(num)
	return output, nil
}

// ConvertToVBucks : returns a raw value as a string
func ConvertToVBucks(num uint) string {
	output := fmt.Sprint("\u24E5", num)
	return output
}
