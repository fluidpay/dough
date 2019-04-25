package currency

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// StringToUint : returns a uint from a string value
func StringToUint(num string) (uint, error) {
	reg := regexp.MustCompile("[0-9]+")
	str := reg.FindAllString(num, -1)
	newStr := strings.Join(str, "")
	output, err := strconv.ParseUint(newStr, 10, 64)
	if err != nil {
		return 0, errorUnableToFormatCurrencyFromString
	}
	return uint(output), nil
}

// DisplayFull : returns a string with full currency formatting... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayFull(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	return formatCurrency(num, ISO), nil
}

// DisplayWithAlpha : returns a string with full currency formatting with the symbol replaced by the ISO three digit alphabetic code... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithAlpha(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := formatCurrency(num, ISO)
	output := swapSymbolWithAlpha(currency, ISO.Symbol, ISO.Alpha)
	return output, nil
}

// DisplayNoSymbol : returns a string with full currency formatting minus the ISO symbol... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayNoSymbol(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := formatCurrency(num, ISO)
	output := removeSymbol(currency, ISO.Symbol)
	return output, nil
}

// DisplayWithDecimal : returns a string with all currency formatting removed except decimal places... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithDecimal(num uint, alpha string) (string, error) {
	ISO, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := formatCurrency(num, ISO)
	currency = removeSymbol(currency, ISO.Symbol)
	return removeDelimiter(currency, ISO.Delimiter), nil
}

// UintToString : returns a string with all currency formatting removed... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func UintToString(num uint, alpha string) (string, error) {
	_, err := getISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	output := fmt.Sprint(num)
	return output, nil
}
