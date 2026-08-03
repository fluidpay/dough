package dough

import (
	"fmt"
	"math"
	"strings"
)

type round string

const (
	Round   round = "round"
	Floor   round = "floor"
	Ceil    round = "ceil"
	Bankers round = "bankers"
)

// currencyByNumeric maps ISO numeric codes to currency entries for O(1) lookup.
var currencyByNumeric map[string]Currency

func init() {
	currencyByNumeric = make(map[string]Currency, len(CurrencyList))
	for _, currency := range CurrencyList {
		currencyByNumeric[currency.Numeric] = currency
	}
}

// GetISOFromNumeric : returns an ISO currency struct or an error if the ISO is not found
func GetISOFromNumeric(num string) (Currency, error) {
	alpha, err := GetAlphaFromISONumeric(num)
	if err != nil {
		return Currency{}, ErrorInvalidISO
	}
	return GetISOFromAlpha(alpha)
}

// GetISOFromAlpha : returns an ISO currency struct or an error if the ISO is not found
func GetISOFromAlpha(alpha string) (Currency, error) {
	currency, ok := CurrencyList[strings.ToUpper(alpha)]
	if !ok {
		return Currency{}, ErrorInvalidISO
	}
	return currency, nil
}

// GetISOCodeFromNumeric : returns a formatted ISO numeric code or an error if the ISO is not found
func GetISOCodeFromNumeric(num string) (string, error) {
	currency, ok := currencyByNumeric[num]
	if !ok {
		return "", ErrorInvalidISO
	}
	return currency.Numeric, nil
}

// GetAlphaFromISONumeric : returns a formatted ISO alpha code from the ISO numeric counterpart
func GetAlphaFromISONumeric(num string) (string, error) {
	currency, ok := currencyByNumeric[num]
	if !ok {
		return "", ErrorInvalidISO
	}
	return currency.Alpha, nil
}

// ConvertToStringWithDecimal : returns the uint as a stringified float
func ConvertToStringWithDecimal(num int, fraction int) string {
	return fmt.Sprintf("%.*f", fraction, float64(float64(num)/math.Pow10(fraction)))
}

// reverseString : returns a reversed string for delimiter formatting
func reverseString(str string) string {
	n := len(str)
	if n < 2 {
		return str
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = str[n-1-i]
	}
	return string(b)
}

// InsertDelimiter : returns a new string with delimiter formatting
func InsertDelimiter(str string, group int, del string) string {
	if group <= 0 || str == "" {
		return str
	}
	var b strings.Builder
	b.Grow(len(str) + len(del)*(len(str)/group))
	for i := 0; i < len(str); i++ {
		if i != 0 && i%group == 0 {
			b.WriteString(del)
		}
		b.WriteByte(str[i])
	}
	return b.String()
}

// SwapSymbolWithAlpha : returns a string with the ISO alpha code instead of symbol
func SwapSymbolWithAlpha(str string, sym string, alpha string) string {
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

// IsNegative : returns a bool based on whether the int is negative or positive
func IsNegative(num int) bool {
	return num < 0
}

// FormatCurrency : returns basic currency formatting
func FormatCurrency(num int, ISO Currency) string {
	isNegativeText := ""
	if num < 0 {
		isNegativeText = "-"
		num = -num
	}

	// to catch frational split panic
	if ISO.Fraction == 0 {
		if !ISO.SymbolPositionFront {
			return fmt.Sprintf("%s%d%s", isNegativeText, num, ISO.Symbol)
		}
		return fmt.Sprintf("%s%s%d", ISO.Symbol, isNegativeText, num)
	}
	str := ConvertToStringWithDecimal(num, ISO.Fraction)
	strSplit := strings.Split(str, ".")
	strSplit[0] = reverseString(strSplit[0])
	strSplit[0] = InsertDelimiter(strSplit[0], ISO.Grouping, ISO.Delimiter)
	strSplit[0] = reverseString(strSplit[0])
	if !ISO.SymbolPositionFront {
		return isNegativeText + strSplit[0] + ISO.Decimal + strSplit[1] + ISO.Symbol
	}
	return ISO.Symbol + isNegativeText + strSplit[0] + ISO.Decimal + strSplit[1]
}

// FloatToInt will take in a float and based upon fraction will output the int version
func FloatToInt(amt float64, fraction int) int {
	return int(math.Round(((amt * 100) * (float64(math.Pow10(fraction)) / 10000)) * 100))
}

// IntToFloat will take in a int and based upon fraction will output the float version
func IntToFloat(amt int, fraction int) float64 {
	return float64(float64(amt) / math.Pow10(fraction))
}

// PercentageFromInt will give you a percentage to the exact precision that you want based on fraction
func PercentageFromInt(amt int, percentage float64, fraction int, round round) float64 {
	// Calculate percentage.
	val := float64(amt) * percentage
	val = val / 100

	// Remove potential rounding errors by moving decimal
	// two places past desired fraction and truncating.
	val = math.Trunc(val*math.Pow10(fraction+2)) / math.Pow10(fraction+2)

	// Handle rounding.
	switch round {
	case Round:
		val = math.Round(val*math.Pow10(fraction)) / math.Pow10(fraction)
	case Floor:
		val = math.Floor(val*math.Pow10(fraction)) / math.Pow10(fraction)
	case Ceil:
		val = math.Ceil(val*math.Pow10(fraction)) / math.Pow10(fraction)
	case Bankers:
		val = math.RoundToEven(val*math.Pow10(fraction)) / math.Pow10(fraction)
	default:
		val = math.Round(val*math.Pow10(fraction)) / math.Pow10(fraction)
	}

	return val
}

// PercentageFromFloat will give you a percentage to the exact precision that you want based on fraction
func PercentageFromFloat(amt float64, percentage float64, fraction int, round round) float64 {
	// Calculate percentage.
	val := amt * percentage
	val = val / 100

	// Remove potential rounding errors by moving decimal
	// two places past desired fraction and truncating.
	val = math.Trunc(val*math.Pow10(fraction+2)) / math.Pow10(fraction+2)

	// Handle rounding.
	switch round {
	case Round:
		val = math.Round(val*math.Pow10(fraction)) / math.Pow10(fraction)
	case Floor:
		val = math.Floor(val*math.Pow10(fraction)) / math.Pow10(fraction)
	case Ceil:
		val = math.Ceil(val*math.Pow10(fraction)) / math.Pow10(fraction)
	case Bankers:
		val = math.RoundToEven(val*math.Pow10(fraction)) / math.Pow10(fraction)
	default:
		val = math.Round(val*math.Pow10(fraction)) / math.Pow10(fraction)
	}

	return val
}
