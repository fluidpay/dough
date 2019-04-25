package currency

import (
	"fmt"
	"math"
	"strings"
)

// validateISOCodeAlpha : returns a formatted ISO alpha code or an error if the ISO is not found
func validateISOCodeAlpha(alpha string) (string, error) {
	alpha = strings.ToUpper(alpha)
	for key := range CurrencyList {
		if key == alpha {
			return alpha, nil
		}
	}
	return "", ErrorInvalidISO
}

// validateISOCodeNumeric : returns a formatted ISO numeric code or an error if the ISO is not found
func validateISOCodeNumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Numeric, nil
		}
	}
	return "", ErrorInvalidISO
}

// getISOFromAlpha : returns the Currency struct for that ISO code alpha
func getISOFromAlpha(alpha string) (Currency, error) {
	alpha, err := validateISOCodeAlpha(alpha)
	if err != nil {
		return Currency{}, err
	}
	return CurrencyList[alpha], nil
}

// getAlphaFromISOCodeNumeric : returns a formatted ISO alpha code from the ISO numeric counterpart
func getAlphaFromISOCodeNumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Alpha, nil
		}
	}
	return "", ErrorInvalidISO
}

// convertToString : returns the uint number as a string
func convertToStringWithDecimal(num uint, exp int) string {
	floatNum := float64(num)
	newNum := float64(floatNum / math.Pow10(exp))
	return fmt.Sprintf("%.*f", exp, newNum)
}

// splitString : returns a map of strings separated by decimal place
func splitString(str string) []string {
	return strings.Split(str, ".")
}

// reverseString : returns a reversed string for delimiter formatting
func reverseString(str string) string {
	var output string
	for key := len(str) - 1; key >= 0; key-- {
		output += string(str[key])
	}
	return output
}

// insertDelimiter : returns a new string with delimiter formatting
func insertDelimiter(str string, group int, del string) string {
	var output strings.Builder
	for key, val := range str {
		if key%group == 0 && key != 0 {
			output.WriteString(del + string(val))
		} else {
			output.WriteString(string(val))
		}
	}
	return output.String()
}

// swapSymbolWithAlpha : returns a string with the ISO alpha code instead of symbol
func swapSymbolWithAlpha(str string, sym string, alpha string) string {
	return strings.Replace(str, sym, alpha+" ", -1)
}

// removeSymbol : returns a string with the symbol removed
func removeSymbol(str string, sym string) string {
	return strings.Replace(str, sym, "", -1)
}

// removeDelimiter : returns a string with the delimiter removed
func removeDelimiter(str string, del string) string {
	return strings.Replace(str, del, "", -1)
}

// removeDecimal : returns a string with the decimal removed
func removeDecimal(str string, dec string) string {
	return strings.Replace(str, dec, "", -1)
}

// formatCurrency : returns basic currency formatting
func formatCurrency(num uint, ISO Currency) string {
	str := convertToStringWithDecimal(num, ISO.Exponent)
	strSplit := splitString(str)
	strStart := strSplit[0]
	strEnd := strSplit[1]
	strReverse := reverseString(strStart)
	strFlipped := insertDelimiter(strReverse, ISO.Grouping, ISO.Delimiter)
	strStart = reverseString(strFlipped)
	return strings.Join([]string{ISO.Symbol + strStart, strEnd}, ISO.Decimal)
}
