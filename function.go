package currency

import (
	"fmt"
	"math"
	"strings"
)

// GetISOFromAlpha : returns a formatted ISO alpha code or an error if the ISO is not found
func GetISOFromAlpha(alpha string) (Currency, error) {
	alpha = strings.ToUpper(alpha)
	for key := range CurrencyList {
		if key == alpha {
			return CurrencyList[key], nil
		}
	}
	return Currency{}, ErrorInvalidISO
}

// GetISOCodeFromNumeric : returns a formatted ISO numeric code or an error if the ISO is not found
func GetISOCodeFromNumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Numeric, nil
		}
	}
	return "", ErrorInvalidISO
}

// GetAlphaFromISONumeric : returns a formatted ISO alpha code from the ISO numeric counterpart
func GetAlphaFromISONumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Alpha, nil
		}
	}
	return "", ErrorInvalidISO
}

// ConvertToStringWithDecimal : returns the uint as a stringified float
func ConvertToStringWithDecimal(num uint, fraction int) string {
	return fmt.Sprintf("%.*f", fraction, float64(float64(num)/math.Pow10(fraction)))
}

// ReverseString : returns a reversed string for delimiter formatting
func ReverseString(str string) string {
	var output string
	for key := len(str) - 1; key >= 0; key-- {
		output += string(str[key])
	}
	return output
}

// InsertDelimiter : returns a new string with delimiter formatting
func InsertDelimiter(str string, group int, del string) string {
	output := ""
	for key, val := range str {
		if key%group == 0 && key != 0 {
			output += del + string(val)
		} else {
			output += string(val)
		}
	}
	return output
}

// SwapSymbolWithAlpha : returns a string with the ISO alpha code instead of symbol
func SwapSymbolWithAlpha(str string, sym string, alpha string) string {
	return strings.Replace(str, sym, alpha+" ", -1)
}

// RemoveSymbol : returns a string with the symbol removed
func RemoveSymbol(str string, sym string) string {
	return strings.Replace(str, sym, "", -1)
}

// RemoveDelimiter : returns a string with the delimiter removed
func RemoveDelimiter(str string, del string) string {
	return strings.Replace(str, del, "", -1)
}

// RemoveDecimal : returns a string with the decimal removed
func RemoveDecimal(str string, dec string) string {
	return strings.Replace(str, dec, "", -1)
}

// FormatCurrency : returns basic currency formatting
func FormatCurrency(num uint, ISO Currency) string {
	str := ConvertToStringWithDecimal(num, ISO.Fraction)
	strSplit := strings.Split(str, ".")
	strSplit[0] = ReverseString(strSplit[0])
	strSplit[0] = InsertDelimiter(strSplit[0], ISO.Grouping, ISO.Delimiter)
	strSplit[0] = ReverseString(strSplit[0])
	if ISO.SymbolPositionFront != true {
		return strSplit[0] + ISO.Decimal + strSplit[1] + ISO.Symbol
	}
	return ISO.Symbol + strSplit[0] + ISO.Decimal + strSplit[1]
}
