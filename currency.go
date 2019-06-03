package dough

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

// StringToInt : returns a int from a string value
func StringToInt(num string, alpha string) (int, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return 0, err
	}

	isNegative := strings.Contains(num, "-")

	reg := regexp.MustCompile("[0-9]+")
	str := reg.FindAllString(num, -1)
	strJoin := strings.Join(str, "")
	fl, err := strconv.ParseFloat(strJoin, 64)

	if err != nil {
		return 0, ErrorUnableToFormatCurrencyFromString
	}
	if strings.Contains(num, ISO.Decimal) == true {
		split := strings.Split(num, ISO.Decimal)
		if len(split[1]) != ISO.Fraction {
			return 0, ErrorUnableToFormatCurrencyFromString
		}
		if isNegative {
			return int(fl) * -1, nil
		}
		return int(fl), nil
	}
	if isNegative {
		return int(fl*math.Pow10(ISO.Fraction)) * -1, nil
	}
	return int(fl * math.Pow10(ISO.Fraction)), nil
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
	return RemoveSymbol(currency, ISO.Symbol), nil
}

// DisplayWithDecimal : returns a string with all currency formatting removed except decimal places... "num" being the amount, "alpha" being the ISO three digit alphabetic code.
func DisplayWithDecimal(num int, alpha string) (string, error) {
	ISO, err := GetISOFromAlpha(alpha)
	if err != nil {
		return "", err
	}
	currency := FormatCurrency(num, ISO)
	currency = RemoveSymbol(currency, ISO.Symbol)
	return RemoveDelimiter(currency, ISO.Delimiter), nil
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
