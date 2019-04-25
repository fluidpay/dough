package currency

import (
	"fmt"
	"math"
	"strings"
)

// ValidateISOCodeAlpha : returns a formatted ISO alpha code or an error if the ISO is not found
func ValidateISOCodeAlpha(alpha string) (string, error) {
	alpha = strings.ToUpper(alpha)
	for key := range CurrencyList {
		if key == alpha {
			return alpha, nil
		}
	}
	return "", ErrorInvalidISO
}

// ValidateISOCodeNumeric : returns a formatted ISO numeric code or an error if the ISO is not found
func ValidateISOCodeNumeric(num string) (string, error) {
	for _, value := range CurrencyList {
		if value.Numeric == num {
			return value.Numeric, nil
		}
	}
	return "", ErrorInvalidISO
}

// GetISOFromAlpha : returns the Currency struct for that ISO code alpha
func GetISOFromAlpha(alpha string) (Currency, error) {
	alpha, err := ValidateISOCodeAlpha(alpha)
	if err != nil {
		return Currency{}, err
	}
	return CurrencyList[alpha], nil
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
func ConvertToStringWithDecimal(num uint, exp int) string {
	return fmt.Sprintf("%.*f", exp, float64(float64(num)/math.Pow10(exp)))
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
	strStart := strSplit[0]
	strEnd := strSplit[1]
	strStart = ReverseString(strStart)
	strStart = InsertDelimiter(strStart, ISO.Grouping, ISO.Delimiter)
	strStart = ReverseString(strStart)
	return ISO.Symbol + strStart + ISO.Decimal + strEnd
}
