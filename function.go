package dough

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type round string

const (
	Round   round = "round"
	Floor   round = "floor"
	Ceil    round = "ceil"
	Bankers round = "bankers"
)

// GetISOFromNumeric : returns an ISO currency struct or an error if the ISO is not found
func GetISOFromNumeric(num string) (Currency, error) {
	alpha, err := GetAlphaFromISONumeric(num)
	if err != nil {
		return Currency{}, ErrorInvalidISO
	}
	currency, err := GetISOFromAlpha(alpha)
	if err != nil {
		return Currency{}, ErrorInvalidISO
	}
	return currency, nil
}

// GetISOFromAlpha : returns an ISO currency struct or an error if the ISO is not found
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

// reverseString : returns a reversed string for delimiter formatting
func reverseString(str string) string {
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
	if math.Signbit(float64(num)) == true {
		return true
	}
	return false
}

// FormatCurrency : returns basic currency formatting
func FormatCurrency(num int, ISO Currency) string {
	isNegative := IsNegative(num)
	isNegativeText := ""
	if isNegative {
		isNegativeText = "-"
	}
	num = int(math.Abs(float64(num)))

	// to catch frational split panic
	if ISO.Fraction == 0 {
		if ISO.SymbolPositionFront != true {
			return fmt.Sprintf("%s%d%s", isNegativeText, num, ISO.Symbol)
		}
		return fmt.Sprintf("%s%s%d", ISO.Symbol, isNegativeText, num)
	}
	str := ConvertToStringWithDecimal(num, ISO.Fraction)
	strSplit := strings.Split(str, ".")
	strSplit[0] = reverseString(strSplit[0])
	strSplit[0] = InsertDelimiter(strSplit[0], ISO.Grouping, ISO.Delimiter)
	strSplit[0] = reverseString(strSplit[0])
	if ISO.SymbolPositionFront != true {
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
	return PercentageFromFloat(float64(amt), percentage, fraction, round)
}

// PercentageFromFloat will give you a percentage to the exact precision that you want based on fraction
func PercentageFromFloat(amt float64, percentage float64, fraction int, round round) float64 {
	// Convert amt to a string.
	amtStr := strconv.FormatFloat(amt, 'f', -1, 64)

	// Get mantissa length.
	var amtMantissaLen int
	split := strings.Split(amtStr, ".")
	if len(split) == 2 {
		amtMantissaLen = len(split[1])
	}

	// Convert percentage to a string.
	percStr := strconv.FormatFloat(percentage, 'f', -1, 64)

	// Get mantissa length.
	var percMantissaLen int
	split = strings.Split(percStr, ".")
	if len(split) == 2 {
		percMantissaLen = len(split[1])
	}

	amtIntStr := strings.Replace(amtStr, ".", "", -1)
	amtInt, _ := strconv.ParseInt(amtIntStr, 10, 64)
	percIntStr := strings.Replace(percStr, ".", "", -1)
	percInt, _ := strconv.ParseInt(percIntStr, 10, 64)

	multi := amtInt * percInt

	// Convert integer back to a string.
	multiStr := strconv.FormatInt(multi, 10)

	// Handle negative sign if present.
	var sign string
	if string(multiStr[0]) == "-" {
		sign = "-"
		multiStr = multiStr[1:]
	}

	// Determine decimal point placement.
	// (1.33 amount = 2, 1.233 percentage = 3, 2 + 3 = 5)
	// (amt mantissa + perc mantissa)
	// Start from end of integer.
	decimalPlace := amtMantissaLen + percMantissaLen

	// Fix potential negative decimal place.
	if len(multiStr)-decimalPlace < 0 {
		neg := int(math.Abs(float64(len(multiStr) - decimalPlace)))

		for i := 0; i < neg; i++ {
			multiStr = "0" + multiStr
		}
	}

	// Add standard percentage of 2 decimal places.
	decimalPlace += 2

	// Fix potential negative decimal place.
	if len(multiStr)-decimalPlace < 0 {
		neg := int(math.Abs(float64(len(multiStr) - decimalPlace)))

		for i := 0; i < neg; i++ {
			multiStr = "0" + multiStr
		}
	}

	// Format into whole and decimal parts.
	var whole string
	var decimal string

	if string(multiStr[0]) == "0" {
		whole = "0"
		decimal = multiStr
	} else {
		whole = multiStr[0 : len(multiStr)-decimalPlace]
		decimal = multiStr[len(multiStr)-decimalPlace:]
	}

	// Form the final result of the calculation.
	multiResultStr := fmt.Sprintf("%v.%v", whole, decimal)

	// If the fraction amount is greater than or equal to the
	// decimal length, return it.
	if fraction >= len(decimal) {
		multiResultStr = sign + multiResultStr
		endNum, _ := strconv.ParseFloat(multiResultStr, 64)
		return endNum
	}

	// If the fraction is 0, round it as is.
	var numToRound float64
	if fraction == 0 {
		multiResultStr = sign + multiResultStr
		numToRound, _ = strconv.ParseFloat(multiResultStr, 64)

		switch round {
		case Round:
			return math.Round(numToRound)
		case Floor:
			return math.Floor(numToRound)
		case Ceil:
			return math.Ceil(numToRound)
		case Bankers:
			return math.RoundToEven(numToRound)
		default:
			return math.Round(numToRound)
		}
	} else {
		// Otherwise, round to given fraction.
		numToRoundStr := fmt.Sprintf("%v%v%v.%v",
			sign,
			whole,
			decimal[0:fraction],
			decimal[fraction:])

		numToRound, _ = strconv.ParseFloat(numToRoundStr, 64)
	}

	var rounded float64
	switch round {
	case Round:
		rounded = math.Round(numToRound)
	case Floor:
		rounded = math.Floor(numToRound)
	case Ceil:
		rounded = math.Ceil(numToRound)
	case Bankers:
		rounded = math.RoundToEven(numToRound)
	default:
		rounded = math.Round(numToRound)
	}

	// Convert rounded num back to string
	roundedStr := strconv.FormatInt(int64(rounded), 10)

	// Handle negative sign if present.
	if string(roundedStr[0]) == "-" {
		sign = "-"
		roundedStr = roundedStr[1:]
	}

	// If rounded string length is less then fraction, pad.
	if len(roundedStr) < fraction {
		for i := 0; i < fraction; i++ {
			roundedStr = "0" + roundedStr
		}
	}

	// Form final whole and decimal numbers.
	whole = roundedStr[0 : len(roundedStr)-fraction]
	decimal = roundedStr[len(roundedStr)-fraction:]

	endNumStr := fmt.Sprintf("%v%v.%v", sign, whole, decimal)
	endNum, _ := strconv.ParseFloat(endNumStr, 64)

	return endNum
}
