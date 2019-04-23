package currency

import "testing"

var TestValidateISOCodeAlphaData = []struct {
	Input  string
	Output string
}{
	{"", errorInvalidISO.Error()},
	{"ABC", errorInvalidISO.Error()},
	{"ABCD", errorInvalidISO.Error()},
	{"usd", "USD"},
}

func TestValidateISOCodeAlpha(t *testing.T) {
	for _, v := range TestValidateISOCodeAlphaData {
		result, err := validateISOCodeAlpha(v.Input)
		if err != nil {
			t.Log(v.Input, "-->", err.Error())
			if err.Error() != v.Output {
				t.Fatal()
			}
		} else {
			if result != v.Output {
				t.Fatal()
			} else {
				t.Log(v.Input, "-->", result)
			}
		}
	}
}

var TestValidateISOCodeNumericData = []struct {
	Input  string
	Output string
}{
	{"", errorInvalidISO.Error()},
	{"123", errorInvalidISO.Error()},
	{"1234", errorInvalidISO.Error()},
	{"840", "840"},
}

func TestValidateISOCodeNumeric(t *testing.T) {
	for _, v := range TestValidateISOCodeNumericData {
		result, err := validateISOCodeNumeric(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Fatal()
			}
			t.Log(v.Input, "-->", err.Error())
		} else {
			if result != v.Output {
				t.Fatal()
			} else {
				t.Log(v.Input, "-->", result)
			}
		}
	}
}

var TestGetISOFromAlphaData = []struct {
	Input  string
	Output interface{}
}{
	{"", errorInvalidISO.Error()},
	{"USA", errorInvalidISO.Error()},
	{"USAA", errorInvalidISO.Error()},
	{"USD", Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "$", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}},
}

func TestGetISOFromAlpha(t *testing.T) {
	for _, v := range TestGetISOFromAlphaData {
		result, err := getISOFromAlpha(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Fatal()
			}
			t.Log(v.Input, "-->", err)
		} else {
			if result != v.Output {
				t.Fatal()
			}
			t.Log(v.Input, "-->", result)
		}
	}
}

var TestGetAlphaFromISOCodeNumericData = []struct {
	Input  string
	Output string
}{
	{"", errorInvalidISO.Error()},
	{"000", errorInvalidISO.Error()},
	{"12345", errorInvalidISO.Error()},
	{"840", "USD"},
}

func TestGetAlphaFromISOCodeNumeric(t *testing.T) {
	for _, v := range TestGetAlphaFromISOCodeNumericData {
		result, err := getAlphaFromISOCodeNumeric(v.Input)
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

var TestConvertToStringData = []struct {
	Num    uint
	Exp    int
	Output string
}{
	{uint(0), 0, "0"},
	{uint(0), 2, "0.00"},
	{uint(1), 2, "0.01"},
	{uint(10), 2, "0.10"},
	{uint(100), 2, "1.00"},
	{uint(1000), 2, "10.00"},
	{uint(10000), 2, "100.00"},
	{uint(100000), 2, "1000.00"},
	{uint(1000000), 2, "10000.00"},
	{uint(10000000), 2, "100000.00"},
	{uint(100000000), 2, "1000000.00"},
}

func TestConvertToString(t *testing.T) {
	for _, v := range TestConvertToStringData {
		result := convertToString(v.Num, v.Exp)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Num, v.Exp, "-->", result)
	}
}

var TestSplitStringData = []struct {
	Input  string
	Output []string
}{
	{"0.00", []string{"0", "00"}},
	{"0.01", []string{"0", "01"}},
	{"0.10", []string{"0", "10"}},
	{"1.00", []string{"1", "00"}},
	{"10.00", []string{"10", "00"}},
	{"100.00", []string{"100", "00"}},
	{"1000.00", []string{"1000", "00"}},
}

func TestSplitString(t *testing.T) {
	for _, v := range TestSplitStringData {
		result := splitString(v.Input)
		if result[0] != v.Output[0] || result[1] != v.Output[1] {
			t.Fatal()
		}
		t.Log(v.Input, "-->", result)
	}
}

var TestReverseStringData = []struct {
	Input  string
	Output string
}{
	{"0", "0"},
	{"1", "1"},
	{"01", "10"},
	{"001", "100"},
	{"000,1", "1,000"},
	{"000,01", "10,000"},
	{"000,001", "100,000"},
	{"000,000,1", "1,000,000"},
}

func TestReverseString(t *testing.T) {
	for _, v := range TestReverseStringData {
		result := reverseString(v.Input)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Input, "-->", result)
	}
}

var TestInsertDelimiterData = []struct {
	Str    string
	Sep    int
	Del    string
	Output string
}{
	{"0", 3, ",", "0"},
	{"01", 3, ",", "01"},
	{"001", 3, ",", "001"},
	{"0001", 3, ",", "000,1"},
	{"00001", 3, ",", "000,01"},
	{"000001", 3, ",", "000,001"},
	{"0000001", 3, ",", "000,000,1"},
}

func TestInsertDelimiter(t *testing.T) {
	for _, v := range TestInsertDelimiterData {
		result := insertDelimiter(v.Str, v.Sep, v.Del)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Str, v.Sep, v.Del, "-->", result)
	}
}

var TestSwapSymbolWithAlphaData = []struct {
	Str    string
	Sym    string
	Alpha  string
	Output string
}{
	{"$0.00", "$", "USD", "USD 0.00"},
	{"$10.00", "$", "USD", "USD 10.00"},
	{"$100.00", "$", "USD", "USD 100.00"},
	{"$1,000.00", "$", "USD", "USD 1,000.00"},
	{"$10,000.00", "$", "USD", "USD 10,000.00"},
	{"$100,000.00", "$", "USD", "USD 100,000.00"},
	{"$1,000,000.00", "$", "USD", "USD 1,000,000.00"},
}

func TestSwapSymbolWithAlpha(t *testing.T) {
	for _, v := range TestSwapSymbolWithAlphaData {
		result := swapSymbolWithAlpha(v.Str, v.Sym, v.Alpha)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Str, v.Sym, v.Alpha, "-->", result)
	}
}

var TestRemoveSymbolData = []struct {
	Str    string
	Sym    string
	Output string
}{
	{"$0.00", "$", "0.00"},
	{"$0.01", "$", "0.01"},
	{"$0.10", "$", "0.10"},
	{"$1.00", "$", "1.00"},
	{"$10.00", "$", "10.00"},
	{"$100.00", "$", "100.00"},
	{"$1,000.00", "$", "1,000.00"},
	{"$10,000.00", "$", "10,000.00"},
	{"$100,000.00", "$", "100,000.00"},
	{"$1,000,000.00", "$", "1,000,000.00"},
}

func TestRemoveSymbol(t *testing.T) {
	for _, v := range TestRemoveSymbolData {
		result := removeSymbol(v.Str, v.Sym)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Str, v.Sym, "-->", result)
	}
}

var TestRemoveDelimiterData = []struct {
	Str    string
	Del    string
	Output string
}{
	{"$0.00", ",", "$0.00"},
	{"$0.01", ",", "$0.01"},
	{"$0.10", ",", "$0.10"},
	{"$1.00", ",", "$1.00"},
	{"$10.00", ",", "$10.00"},
	{"$100.00", ",", "$100.00"},
	{"$1,000.00", ",", "$1000.00"},
	{"$10,000.00", ",", "$10000.00"},
	{"$100,000.00", ",", "$100000.00"},
	{"$1,000,000.00", ",", "$1000000.00"},
}

func TestRemoveDelimiter(t *testing.T) {
	for _, v := range TestRemoveDelimiterData {
		result := removeDelimiter(v.Str, v.Del)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Str, v.Del, "-->", result)
	}
}

var TestRemoveDecimalData = []struct {
	Str    string
	Dec    string
	Output string
}{
	{"$0.00", ".", "$000"},
	{"$0.01", ".", "$001"},
	{"$0.10", ".", "$010"},
	{"$1.00", ".", "$100"},
	{"$10.00", ".", "$1000"},
	{"$100.00", ".", "$10000"},
	{"$1,000.00", ".", "$1,00000"},
	{"$10,000.00", ".", "$10,00000"},
	{"$100,000.00", ".", "$100,00000"},
	{"$1,000,000.00", ".", "$1,000,00000"},
}

func TestRemoveDecimal(t *testing.T) {
	for _, v := range TestRemoveDecimalData {
		result := removeDecimal(v.Str, v.Dec)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Str, v.Dec, "-->", result)
	}
}

var TestFormatCurrencyData = []struct {
	Num    uint
	ISO    Currency
	Output string
}{
	{uint(0), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$0.00"},
	{uint(1), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$0.01"},
	{uint(10), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$0.10"},
	{uint(100), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$1.00"},
	{uint(1000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$10.00"},
	{uint(10000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$100.00"},
	{uint(100000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$1,000.00"},
	{uint(1000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$10,000.00"},
	{uint(10000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$100,000.00"},
	{uint(100000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}, "$1,000,000.00"},
}

func TestFormatCurrency(t *testing.T) {
	for _, v := range TestFormatCurrencyData {
		result := formatCurrency(v.Num, v.ISO)
		if result != v.Output {
			t.Fatal()
		}
		t.Log(v.Num, "-->", result)
	}
}
