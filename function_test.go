package dough

import "testing"

var TestValidateISOCodeAlphaData = []struct {
	Input  string
	Output string
}{
	{"", ErrorInvalidISO.Error()},
	{"ABC", ErrorInvalidISO.Error()},
	{"ABCD", ErrorInvalidISO.Error()},
	{"usd", "USD"},
}

func TestValidateISOCodeAlpha(t *testing.T) {
	for _, v := range TestValidateISOCodeAlphaData {
		curr, err := GetISOFromAlpha(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err)
			}
		} else if curr.Alpha != v.Output {
			t.Error(curr.Alpha)
		}
	}
}

var TestValidateISOCodeNumericData = []struct {
	Input  string
	Output string
}{
	{"", ErrorInvalidISO.Error()},
	{"123", ErrorInvalidISO.Error()},
	{"1234", ErrorInvalidISO.Error()},
	{"840", "840"},
}

func TestValidateISOCodeNumeric(t *testing.T) {
	for _, v := range TestValidateISOCodeNumericData {
		result, err := GetISOCodeFromNumeric(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err)
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestGetISOFromAlphaData = []struct {
	Input  string
	Output interface{}
}{
	{"", ErrorInvalidISO.Error()},
	{"USA", ErrorInvalidISO.Error()},
	{"USAA", ErrorInvalidISO.Error()},
	{"USD", Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "$", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}},
}

func TestGetISOFromAlpha(t *testing.T) {
	for _, v := range TestGetISOFromAlphaData {
		result, err := GetISOFromAlpha(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err)
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestGetAlphaFromISOCodeNumericData = []struct {
	Input  string
	Output string
}{
	{"", ErrorInvalidISO.Error()},
	{"000", ErrorInvalidISO.Error()},
	{"12345", ErrorInvalidISO.Error()},
	{"840", "USD"},
}

func TestGetAlphaFromISOCodeNumeric(t *testing.T) {
	for _, v := range TestGetAlphaFromISOCodeNumericData {
		result, err := GetAlphaFromISONumeric(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err)
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

var TestConvertToStringWithDecimalData = []struct {
	Num    int
	Exp    int
	Output string
}{
	{int(0), 0, "0"},
	{int(0), 2, "0.00"},
	{int(1), 2, "0.01"},
	{int(10), 2, "0.10"},
	{int(100), 2, "1.00"},
	{int(1000), 2, "10.00"},
	{int(10000), 2, "100.00"},
	{int(100000), 2, "1000.00"},
	{int(1000000), 2, "10000.00"},
	{int(10000000), 2, "100000.00"},
	{int(100000000), 2, "1000000.00"},
}

func TestConvertToStringWithDecimal(t *testing.T) {
	for _, v := range TestConvertToStringWithDecimalData {
		result := ConvertToStringWithDecimal(v.Num, v.Exp)
		if result != v.Output {
			t.Error(result)
		}
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
		result := ReverseString(v.Input)
		if result != v.Output {
			t.Error(result)
		}
	}
}

var TestInsertDelimiterData = []struct {
	Str    string
	Group  int
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
		result := InsertDelimiter(v.Str, v.Group, v.Del)
		if result != v.Output {
			t.Error(result)
		}
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
		result := SwapSymbolWithAlpha(v.Str, v.Sym, v.Alpha)
		if result != v.Output {
			t.Error(result)
		}
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
		result := RemoveSymbol(v.Str, v.Sym)
		if result != v.Output {
			t.Error(result)
		}
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
		result := RemoveDelimiter(v.Str, v.Del)
		if result != v.Output {
			t.Error(result)
		}
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
		result := RemoveDecimal(v.Str, v.Dec)
		if result != v.Output {
			t.Error(result)
		}
	}
}

var TestIsNegativeData = []struct {
	Num    int
	Output bool
}{
	{int(0), false},
	{int(1), false},
	{int(-1), true},
}

func TestIsNegative(t *testing.T) {
	for _, v := range TestIsNegativeData {
		result := IsNegative(v.Num)
		if result != v.Output {
			t.Error(result)
		}
	}
}

var TestFormatCurrencyData = []struct {
	Num    int
	ISO    Currency
	Output string
}{
	{int(0), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$0.00"},
	{int(1), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$0.01"},
	{int(10), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$0.10"},
	{int(100), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$1.00"},
	{int(1000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$10.00"},
	{int(10000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$100.00"},
	{int(100000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$1,000.00"},
	{int(1000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$10,000.00"},
	{int(10000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$100,000.00"},
	{int(100000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$1,000,000.00"},
	{int(-0), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$0.00"},
	{int(-1), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-0.01"},
	{int(-10), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-0.10"},
	{int(-100), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-1.00"},
	{int(-1000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-10.00"},
	{int(-10000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-100.00"},
	{int(-100000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-1,000.00"},
	{int(-1000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-10,000.00"},
	{int(-10000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-100,000.00"},
	{int(-100000000), Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}, "$-1,000,000.00"},
	{int(0), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "0.00" + "\u0625\u002E\u062F"},
	{int(1), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "0.01" + "\u0625\u002E\u062F"},
	{int(10), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "0.10" + "\u0625\u002E\u062F"},
	{int(100), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "1.00" + "\u0625\u002E\u062F"},
	{int(1000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "10.00" + "\u0625\u002E\u062F"},
	{int(10000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "100.00" + "\u0625\u002E\u062F"},
	{int(100000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "1,000.00" + "\u0625\u002E\u062F"},
	{int(1000000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "10,000.00" + "\u0625\u002E\u062F"},
	{int(10000000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "100,000.00" + "\u0625\u002E\u062F"},
	{int(100000000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "1,000,000.00" + "\u0625\u002E\u062F"},
	{int(-0), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "0.00" + "\u0625\u002E\u062F"},
	{int(-1), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-0.01" + "\u0625\u002E\u062F"},
	{int(-10), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-0.10" + "\u0625\u002E\u062F"},
	{int(-100), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-1.00" + "\u0625\u002E\u062F"},
	{int(-1000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-10.00" + "\u0625\u002E\u062F"},
	{int(-10000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-100.00" + "\u0625\u002E\u062F"},
	{int(-100000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-1,000.00" + "\u0625\u002E\u062F"},
	{int(-1000000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-10,000.00" + "\u0625\u002E\u062F"},
	{int(-10000000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-100,000.00" + "\u0625\u002E\u062F"},
	{int(-100000000), Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: false}, "-1,000,000.00" + "\u0625\u002E\u062F"},
}

func TestFormatCurrency(t *testing.T) {
	for _, v := range TestFormatCurrencyData {
		result := FormatCurrency(v.Num, v.ISO)
		if result != v.Output {
			t.Error(result)
		}
	}
}

func TestFloatToInt(t *testing.T) {
	for _, v := range TestLargeNums {
		result := FloatToInt(v.Float1, 1)
		if result != v.Integer {
			t.Error("Expected: ", v.Integer, "Got: ", result)
		}
		result = FloatToInt(v.Float2, 2)
		if result != v.Integer {
			t.Error("Expected:", v.Integer, "Got: ", result)
		}
		result = FloatToInt(v.Float3, 3)
		if result != v.Integer {
			t.Error("Expected: ", v.Integer, "Got: ", result)
		}
	}
}

func BenchmarkFloatToInt(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FloatToInt(123456789.99, 2)
	}
}

func TestIntToFloat(t *testing.T) {
	for _, v := range TestLargeNums {
		result := IntToFloat(v.Integer, 1)
		if result != v.Float1 {
			t.Error("Expected: ", v.Float1, "Got: ", result)
		}
		result = IntToFloat(v.Integer, 2)
		if result != v.Float2 {
			t.Error("Expected: ", v.Float2, "Got: ", result)
		}
		result = IntToFloat(v.Integer, 3)
		if result != v.Float3 {
			t.Error("Expected: ", v.Float3, "Got: ", result)
		}
	}
}

func BenchmarkIntToFloat(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		IntToFloat(123456789, 2)
	}
}

func TestGetPercentage(t *testing.T) {
	for _, v := range TestLargeNums {
		result := Percentage(v.Integer, .01, 2)
		if result != v.Float2 {
			t.Error("Expected: ", v.Float2, "Got: ", result)
		}
	}
}
