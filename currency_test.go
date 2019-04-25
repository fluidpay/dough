package currency

import "testing"

var TestStringToUintData = []struct {
	Input  string
	Output interface{}
}{
	{"", errorUnableToFormatCurrencyFromString.Error()},
	{"abcd", errorUnableToFormatCurrencyFromString.Error()},
	{"$0.00", uint(0)},
	{"$0.01", uint(1)},
	{"$0.10", uint(10)},
	{"$1.00", uint(100)},
	{"$10.00", uint(1000)},
	{"$100.00", uint(10000)},
	{"$1,000.00", uint(100000)},
	{"$10,000.00", uint(1000000)},
	{"$100,000.00", uint(10000000)},
	{"$1,000,000.00", uint(100000000)},
}

func TestStringToUint(t *testing.T) {
	for _, v := range TestStringToUintData {
		result, err := StringToUint(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err.Error())
			}
		} else {
			if result != v.Output {
				t.Error(result)
			}
			t.Log(v.Input, "-->", result)
		}
	}
}

var TestDisplayFullData = []struct {
	Amount uint
	Alpha  string
	Output string
}{
	{0, "USA", errorInvalidISO.Error()},
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
				t.Fatal()
			} else {
				t.Log(v.Amount, v.Alpha, "-->", err.Error())
			}
		} else {
			if result != v.Output {
				t.Fatal()
			} else {
				t.Log(v.Amount, v.Alpha, "-->", result)
			}
		}
	}
}

var TestDisplayWithAlphaData = []struct {
	Amount uint
	Alpha  string
	Output string
}{
	{0, "USA", errorInvalidISO.Error()},
	{0, "USD", "USD 0.00"},
	// {0, "AED", "AED 0.00"},
	// {0, "ARS", "ARS 0,00"},
	// {0, "AUD", "AUD 0.00"},
}

func TestDisplayWithAlpha(t *testing.T) {
	for _, v := range TestDisplayWithAlphaData {
		result, err := DisplayWithAlpha(v.Amount, v.Alpha)
		if err != nil {
			if err.Error() != v.Output {
				t.Fatal()
			} else {
				t.Log(v.Amount, v.Alpha, "-->", err.Error())
			}
		} else {
			if result != v.Output {
				t.Fatal()
			} else {
				t.Log(v.Amount, v.Alpha, "-->", result)
			}
		}
	}
}

var TestDisplayNoSymbolData = []struct {
	Num    uint
	Alpha  string
	Output string
}{
	{0, "USA", errorInvalidISO.Error()},
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
			t.Log(err.Error())
			if err.Error() != v.Output {
				t.Fatal()
			}
		} else {
			if result != v.Output {
				t.Fatal()
			}
			t.Log(v.Num, v.Alpha, "-->", result)
		}
	}
}

var TestDisplayWithDecimalData = []struct {
	Num    uint
	Alpha  string
	Output string
}{
	{0, "USA", errorInvalidISO.Error()},
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
			t.Log(err.Error())
			if err.Error() != v.Output {
				t.Fatal()
			}
		} else {
			if result != v.Output {
				t.Fatal()
			}
			t.Log(v.Num, v.Alpha, "-->", result)
		}
	}
}

var TestUintToStringData = []struct {
	Num    uint
	Alpha  string
	Output string
}{
	{0, "USA", errorInvalidISO.Error()},
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
			t.Log(err.Error())
			if err.Error() != v.Output {
				t.Fatal()
			}
		} else {
			if result != v.Output {
				t.Fatal()
			}
			t.Log(v.Num, v.Alpha, "-->", result)
		}
	}
}
