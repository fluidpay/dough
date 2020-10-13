package dough

import (
	"regexp"
	"strconv"
	"strings"
)

// StringToInt : returns a int from a string value
func StringToInt(num string, alpha string, options ...bool) (int, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return 0, err
	}

	// Clean string
	reg, _ := regexp.Compile("[^-" + ISO.Decimal + "0-9]+")
	str := reg.ReplaceAllString(num, "")
	str = strings.Replace(str, ISO.Decimal, ".", -1) // Replace ISO specific decimal with float decimal .
	if str == "" {
		return 0, ErrorInvalidStringFormat
	}

	// Validate ISO fraction matches
	split := strings.Split(str, ".")

	// Check valid fraction match
	allowLoose := false
	if len(options) >= 1 {
		allowLoose = options[0]
	}
	if ISO.Fraction != 0 {
		if !allowLoose && len(split) == 2 && len(split[1]) != ISO.Fraction {
			return 0, ErrorInvalidISOFractionMatch
		}

		// Convert to Float - to test if valid number
		fl, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return 0, ErrorInvalidStringFormat
		}

		// Convert float to cents based upon iso fraction
		return FloatToInt(fl, ISO.Fraction), nil
	}

	s := strings.Replace(str, ".", "", -1)
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, ErrorInvalidStringFormat
	}
	return i, nil
}

// DisplayFull : returns a string with full currency formatting... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayFull(num int, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	return FormatCurrency(num, ISO), nil
}

// DisplayWithAlpha : returns a string with full currency formatting with the symbol replaced by the ISO three digit alphabetic code... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithAlpha(num int, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	return SwapSymbolWithAlpha(currency, ISO.Symbol, ISO.Alpha), nil
}

// DisplayNoSymbol : returns a string with full currency formatting minus the ISO symbol... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayNoSymbol(num int, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	return removeSymbol(currency, ISO.Symbol), nil
}

// DisplayWithDecimal : returns a string with all currency formatting removed except decimal places... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithDecimal(num int, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	currency = removeSymbol(currency, ISO.Symbol)
	return removeDelimiter(currency, ISO.Delimiter), nil
}

// TopCurrencies returns the list of top currencies based upon usage
func TopCurrencies() ([]Currency, error) {
	return ListCurrencies([]string{"USD", "EUR", "GBP", "INR", "CRC", "VND", "HUF", "ILS", "CNY", "KRW", "NGN", "PYG", "PHP", "PLN", "THB", "UAH", "JPY"})
}

// ListCurrencies : returns a list of currencies
func ListCurrencies(list []string) ([]Currency, error) {
	currencies := []Currency{}
	for _, v := range list {
		ISO, err := GetISOFromAlpha(v)
		if err != nil {
			return nil, err
		}
		currencies = append(currencies, ISO)
	}
	return currencies, nil
}
