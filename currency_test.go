package dough

import (
	"reflect"
	"testing"
)

func TestCurrencyCount(t *testing.T) {
	num := len(CurrencyList)
	t.Log("Currency Count: ", num)
}

var TestStringToIntData = []struct {
	Num    string
	Alpha  string
	Output interface{}
}{
	{"", "USD", ErrorUnableToFormatCurrencyFromString.Error()},
	{"     ", "USD", ErrorUnableToFormatCurrencyFromString.Error()},
	{"abcd", "USD", ErrorUnableToFormatCurrencyFromString.Error()},
	{"$5", "USA", ErrorInvalidISO.Error()},
	{"$5", "USD", 500},
	{"$500", "USD", 50000},
	{"$-500", "USD", -50000},
	{"$05", "USD", 500},
	{"$0.05", "USD", 5},
	{"$5.0", "USD", ErrorInvalidStringFormat.Error()},
	{"$5.000", "USD", ErrorInvalidStringFormat.Error()},
	{"$5.52", "USD", 552},
	{"$0.00", "USD", 0},
	{"$0.01", "USD", 1},
	{"$0.10", "USD", 10},
	{"$1.00", "USD", 100},
	{"$10.00", "USD", 1000},
	{"$100.00", "USD", 10000},
	{"$1,000.00", "USD", 100000},
	{"$10,000.00", "USD", 1000000},
	{"$100,000.00", "USD", 10000000},
	{"$1,000,000.00", "USD", 100000000},

	// Problematic Numbers
	{"$538.92", "USD", 53892},
	{"$65.85", "USD", 6585},
	{"$17.99", "USD", 1799},
	{"538.92", "USD", 53892},
	{"65.85", "USD", 6585},
	{"17.99", "USD", 1799},
	{"$-538.92", "USD", -53892},
	{"$-65.85", "USD", -6585},
	{"$-17.99", "USD", -1799},
	{"-538.92", "USD", -53892},
	{"-65.85", "USD", -6585},
	{"-17.99", "USD", -1799},
	{"-$538.92", "USD", -53892},
	{"-$65.85", "USD", -6585},
	{"-$17.99", "USD", -1799},

	// Non USD
	{"$100.00,00", "ARS", 1000000},
	{"$10,000,000", "JPY", 10000000},
}

func TestStringToInt(t *testing.T) {
	for _, v := range TestStringToIntData {
		result, err := StringToInt(v.Num, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayFullData = []struct {
	Amount int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "$0.00"},
	{1, "USD", "$0.01"},
	{10, "USD", "$0.10"},
	{100, "USD", "$1.00"},
	{1000, "USD", "$10.00"},
	{10000, "USD", "$100.00"},
	{100000, "USD", "$1,000.00"},
	{1000000, "USD", "$10,000.00"},
	{10000000, "USD", "$100,000.00"},
	{100000000, "USD", "$1,000,000.00"},
	{0, "AED", "0.00\u0625\u002E\u062F"},
	{1, "AED", "0.01\u0625\u002E\u062F"},
	{10, "AED", "0.10\u0625\u002E\u062F"},
	{100, "AED", "1.00\u0625\u002E\u062F"},
	{1000, "AED", "10.00\u0625\u002E\u062F"},
	{10000, "AED", "100.00\u0625\u002E\u062F"},
	{100000, "AED", "1,000.00\u0625\u002E\u062F"},
	{1000000, "AED", "10,000.00\u0625\u002E\u062F"},
	{10000000, "AED", "100,000.00\u0625\u002E\u062F"},
	{100000000, "AED", "1,000,000.00\u0625\u002E\u062F"},
	{-0, "USD", "$0.00"},
	{-1, "USD", "$-0.01"},
	{-10, "USD", "$-0.10"},
	{-100, "USD", "$-1.00"},
	{-1000, "USD", "$-10.00"},
	{-10000, "USD", "$-100.00"},
	{-100000, "USD", "$-1,000.00"},
	{-1000000, "USD", "$-10,000.00"},
	{-10000000, "USD", "$-100,000.00"},
	{-100000000, "USD", "$-1,000,000.00"},
}

func TestDisplayFull(t *testing.T) {
	for _, v := range TestDisplayFullData {
		result, err := DisplayFull(v.Amount, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayWithAlphaData = []struct {
	Amount int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "USD 0.00"},
	{-1, "USD", "USD -0.01"},
}

func TestDisplayWithAlpha(t *testing.T) {
	for _, v := range TestDisplayWithAlphaData {
		result, err := DisplayWithAlpha(v.Amount, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayNoSymbolData = []struct {
	Num    int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "0.00"},
	{1, "USD", "0.01"},
	{10, "USD", "0.10"},
	{100, "USD", "1.00"},
	{1000, "USD", "10.00"},
	{10000, "USD", "100.00"},
	{100000, "USD", "1,000.00"},
	{-1, "USD", "-0.01"},
	{-10, "USD", "-0.10"},
	{-100, "USD", "-1.00"},
	{-1000, "USD", "-10.00"},
	{-10000, "USD", "-100.00"},
	{-100000, "USD", "-1,000.00"},
}

func TestDisplayNoSymbol(t *testing.T) {
	for _, v := range TestDisplayNoSymbolData {
		result, err := DisplayNoSymbol(v.Num, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestDisplayWithDecimalData = []struct {
	Num    int
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "0.00"},
	{1, "USD", "0.01"},
	{10, "USD", "0.10"},
	{100, "USD", "1.00"},
	{1000, "USD", "10.00"},
	{10000, "USD", "100.00"},
	{100000, "USD", "1000.00"},
	{-1, "USD", "-0.01"},
	{-10, "USD", "-0.10"},
	{-100, "USD", "-1.00"},
	{-1000, "USD", "-10.00"},
	{-10000, "USD", "-100.00"},
	{-100000, "USD", "-1000.00"},
}

func TestDisplayWithDecimal(t *testing.T) {
	for _, v := range TestDisplayWithDecimalData {
		result, err := DisplayWithDecimal(v.Num, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

func TestTopCurrencies(t *testing.T) {
	_, err := TopCurrencies()
	if err != nil {
		t.Error("Could not find all valid currencies from list")
	}
}

var TestListCurrenciesData = []struct {
	Input  []string
	Output interface{}
}{
	{[]string{"USA"}, ErrorInvalidISO},
	{[]string{"USD"}, []Currency{{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}}},
	{[]string{"AED"}, []Currency{{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}}},
	{[]string{"USD", "GBP"}, []Currency{{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, {Unit: "Pound Sterling", Alpha: "GBP", Numeric: "826", Symbol: "£", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}}},
}

func TestListCurrencies(t *testing.T) {
	for _, v := range TestListCurrenciesData {
		result, err := ListCurrencies(v.Input)
		if err != nil {
			if err != v.Output {
				t.Error(err.Error())
			}
		} else if reflect.DeepEqual(result, v.Output) != true {
			t.Error(result)
		}
	}
}
