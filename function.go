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
	return "", errorInvalidISO
}

// validateISOCodeNumeric : returns a formatted ISO numeric code or an error if the ISO is not found
func validateISOCodeNumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Numeric, nil
		}
	}
	return "", errorInvalidISO
}

// getISOFromAlpha : returns the Currency struct for that ISO code alpha
func getISOFromAlpha(alpha string) (Currency, error) {
	alpha, err := validateISOCodeAlpha(alpha)
	if err != nil {
		return Currency{}, err
	}
	output := CurrencyList[alpha]
	return output, nil
}

// getAlphaFromISOCodeNumeric : returns a formatted ISO alpha code from the ISO numeric counterpart
func getAlphaFromISOCodeNumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Alpha, nil
		}
	}
	return "", errorInvalidISO
}

// convertToString : returns the uint number as a string
func convertToString(num uint, exp int) string {
	floatNum := float64(num)
	newNum := float64(floatNum / math.Pow10(exp))
	output := fmt.Sprintf("%.*f", exp, newNum)
	return output
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
func insertDelimiter(str string, sep int, del string) string {
	var output strings.Builder
	for key, val := range str {
		if key%sep == 0 && key != 0 {
			fmt.Fprintf(&output, del+string(val))
		} else {
			fmt.Fprintf(&output, string(val))
		}
	}
	output.WriteString("")
	return output.String()
}

// swapSymbolWithAlpha : returns a string with the ISO alpha code instead of symbol
func swapSymbolWithAlpha(str string, sym string, alpha string) string {
	output := strings.Replace(str, sym, alpha+" ", -1)
	return output
}

// removeSymbol : returns a string with the symbol removed
func removeSymbol(str string, sym string) string {
	output := strings.Replace(str, sym, "", -1)
	return output
}

// removeDelimiter : returns a string with the delimiter removed
func removeDelimiter(str string, del string) string {
	output := strings.Replace(str, del, "", -1)
	return output
}

// removeDecimal : returns a string with the decimal removed
func removeDecimal(str string, dec string) string {
	output := strings.Replace(str, dec, "", -1)
	return output
}

// formatCurrency : returns basic currency formatting
func formatCurrency(num uint, ISO Currency) string {
	exp := ISO.Exponent
	dec := ISO.Decimal
	sep := ISO.Separator
	del := ISO.Delimiter
	sym := ISO.Symbol
	str := convertToString(num, exp)
	strSplit := splitString(str)
	strStart := strSplit[0]
	strEnd := strSplit[1]
	strReverse := reverseString(strStart)
	strFlipped := insertDelimiter(strReverse, sep, del)
	strStart = reverseString(strFlipped)
	output := strings.Join([]string{sym + strStart, strEnd}, dec)
	return output
}
