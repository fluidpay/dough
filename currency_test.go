package currency

import "testing"

var TestConvertValueFromStringData = []struct {
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

func TestConvertValueFromString(t *testing.T) {
	for _, v := range TestConvertValueFromStringData {
		result, err := ConvertValueFromString(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Fatal()
			}
			t.Log(v.Input, "-->", err.Error())
		} else {
			if result != v.Output {
				t.Fatal()
			}
			t.Log(v.Input, "-->", result)
		}
	}
}

var TestConvertToStringFullData = []struct {
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

func TestConvertToStringFull(t *testing.T) {
	for _, v := range TestConvertToStringFullData {
		result, err := ConvertToStringFull(v.Amount, v.Alpha)
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

var TestConvertToStringAlphaData = []struct {
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

func TestConvertToStringAlpha(t *testing.T) {
	for _, v := range TestConvertToStringAlphaData {
		result, err := ConvertToStringAlpha(v.Amount, v.Alpha)
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

var TestConvertToStringNoSymbolData = []struct {
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

func TestConvertToStringNoSymbol(t *testing.T) {
	for _, v := range TestConvertToStringNoSymbolData {
		result, err := ConvertToStringNoSymbol(v.Num, v.Alpha)
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

var TestConvertToStringDecimalData = []struct {
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

func TestConvertToStringDecimal(t *testing.T) {
	for _, v := range TestConvertToStringDecimalData {
		result, err := ConvertToStringDecimal(v.Num, v.Alpha)
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

var TestConvertToStringRawData = []struct {
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

func TestConvertToStringRaw(t *testing.T) {
	for _, v := range TestConvertToStringRawData {
		result, err := ConvertToStringRaw(v.Num, v.Alpha)
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

var TestConvertToVBucksData = []struct {
	Input  uint
	Output string
}{
	{uint(0), "\u24E50"},
	{uint(1), "\u24E51"},
}

func TestConvertToVBucks(t *testing.T) {
	for _, v := range TestConvertToVBucksData {
		result := ConvertToVBucks(v.Input)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Input, "-->", result)
	}
}
