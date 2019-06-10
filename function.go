package dough

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
func ConvertToStringWithDecimal(num int, fraction int) string {
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

// IsNegative : returns a bool based on whether the int is negative or positive
func IsNegative(num int) bool {
	if math.Signbit(float64(num)) == true {
		return true
	}
	return false
}

// FormatCurrency : returns basic currency formatting
func FormatCurrency(num int, ISO Currency) string {
	isNegative := IsNegative(num)
	num = int(math.Abs(float64(num)))
	str := ConvertToStringWithDecimal(num, ISO.Fraction)
	strSplit := strings.Split(str, ".")
	strSplit[0] = ReverseString(strSplit[0])
	strSplit[0] = InsertDelimiter(strSplit[0], ISO.Grouping, ISO.Delimiter)
	strSplit[0] = ReverseString(strSplit[0])
	if ISO.SymbolPositionFront != true {
		if isNegative {
			return "-" + strSplit[0] + ISO.Decimal + strSplit[1] + ISO.Symbol
		}
		return strSplit[0] + ISO.Decimal + strSplit[1] + ISO.Symbol
	}
	if isNegative {
		return ISO.Symbol + "-" + strSplit[0] + ISO.Decimal + strSplit[1]
	}
	return ISO.Symbol + strSplit[0] + ISO.Decimal + strSplit[1]
}

// FloatToInt will take in a float and based upon fraction will output the int version
func FloatToInt(amt float64, fraction int) int {
	return int(math.Round(((amt * 100) * (float64(math.Pow10(fraction)) / 10000)) * 100))
}

// IntToFloat will take in a int and based upon magnitude will output the float version
func IntToFloat(amt int, magnitude int) float64 {
	return float64(float64(amt) / math.Pow10(magnitude))
}

// Percentage will give you a percentage to the exact precision that you want based on magnitude
func Percentage(amt int, pct float64, magnitude int) float64 {
	return math.Round(float64(amt)*pct*math.Pow10(magnitude)) / math.Pow10(magnitude)
}
