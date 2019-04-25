package currency

import "testing"

var TestStringToUintData = []struct {
	Num    string
	Alpha  string
	Output interface{}
}{
	{"", "USD", ErrorUnableToFormatCurrencyFromString.Error()},
	{"abcd", "USD", ErrorUnableToFormatCurrencyFromString.Error()},
	{"$5", "USD", uint(500)},
	{"$500", "USD", uint(50000)},
	{"$05", "USD", uint(500)},
	{"$0.05", "USD", uint(5)},
	{"$5.0", "USD", uint(500)},
	{"$0.00", "USD", uint(0)},
	{"$0.01", "USD", uint(1)},
	{"$0.10", "USD", uint(10)},
	{"$1.00", "USD", uint(100)},
	{"$10.00", "USD", uint(1000)},
	{"$100.00", "USD", uint(10000)},
	{"$1,000.00", "USD", uint(100000)},
	{"$10,000.00", "USD", uint(1000000)},
	{"$100,000.00", "USD", uint(10000000)},
	{"$1,000,000.00", "USD", uint(100000000)},
}

func TestStringToUint(t *testing.T) {
	for _, v := range TestStringToUintData {
		result, err := StringToUint(v.Num, v.Alpha)
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
	Amount uint
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
	Amount uint
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "USD 0.00"},
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
	Num    uint
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
	Num    uint
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

var TestUintToStringData = []struct {
	Num    uint
	Alpha  string
	Output string
}{
	{0, "USA", ErrorInvalidISO.Error()},
	{0, "USD", "0"},
	{1, "USD", "1"},
	{10, "USD", "10"},
	{100, "USD", "100"},
	{1000, "USD", "1000"},
	{10000, "USD", "10000"},
	{100000, "USD", "100000"},
}

func TestUintToString(t *testing.T) {
	for _, v := range TestUintToStringData {
		result, err := UintToString(v.Num, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}
