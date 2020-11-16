package dough

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var TestGetISOFromNumericData = []struct {
	Input  string
	Output interface{}
}{
	{"", ErrorInvalidISO.Error()},
	{"USA", ErrorInvalidISO.Error()},
	{"840", Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "$", Fraction: 2, Decimal: ".", Grouping: 3, Delimiter: ",", SymbolPositionFront: true}},
}

func TestGetISOFromNumeric(t *testing.T) {
	for _, v := range TestGetISOFromNumericData {
		result, err := GetISOFromNumeric(v.Input)
		if err != nil {
			if err.Error() != v.Output {
				t.Error(err)
			}
		} else if result != v.Output {
			t.Error(result)
		}
	}
}

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

var TestreverseStringData = []struct {
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
	for _, v := range TestreverseStringData {
		result := reverseString(v.Input)
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

var TestremoveSymbolData = []struct {
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
	for _, v := range TestremoveSymbolData {
		result := removeSymbol(v.Str, v.Sym)
		if result != v.Output {
			t.Error(result)
		}
	}
}

var TestremoveDelimiterData = []struct {
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
	for _, v := range TestremoveDelimiterData {
		result := removeDelimiter(v.Str, v.Del)
		if result != v.Output {
			t.Error(result)
		}
	}
}

var TestremoveDecimalData = []struct {
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
	for _, v := range TestremoveDecimalData {
		result := removeDecimal(v.Str, v.Dec)
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

var intPercentageData = []struct {
	amt      int
	pct      float64
	fraction int
	round    round
	result   float64
}{
	{898, 56.7, 2, Round, 509.17},
	{898, 56.7, 3, Round, 509.166},
	{10975, 11, 2, Round, 1207.25},
	{10975, 11, 3, Round, 1207.25},
	{10, 27, 2, Round, 2.7},
	{6942, 99, 2, Round, 6872.58},
	{9999, 0, 2, Round, 0},
	{9999, 100, 2, Round, 9999},
	{45435, 69, 2, Round, 31350.15},
	{420, 42, 2, Round, 176.4},
	{89357, 85.4, 2, Round, 76310.88},
	{200, 14, 2, Round, 28},
	{1414104958, 21.45, 2, Round, 303325513.49},
	{9857, 75.6, 1, Round, 7451.9},
	{9857, 75.6, 0, Round, 7452},
	{100, 10, 2, Round, 10},
	{1055, 350, 2, Round, 3692.5},
	{100, 1598, 2, Round, 1598},
	{333, 8, 2, Round, 26.64},
	{1992, .345, 2, Round, 6.87},
	{65, .011, 3, Round, 0.007},
	{65, .011, 4, Round, 0.0072},

	{898, 56.7, 2, Floor, 509.16},
	{898, 56.7, 3, Floor, 509.166},
	{10975, 11, 2, Floor, 1207.25},
	{10975, 11, 3, Floor, 1207.25},
	{10, 27, 2, Floor, 2.7},
	{6942, 99, 2, Floor, 6872.58},
	{9999, 0, 2, Floor, 0},
	{9999, 100, 2, Floor, 9999},
	{45435, 69, 2, Floor, 31350.15},
	{420, 42, 2, Floor, 176.4},
	{89357, 85.4, 2, Floor, 76310.87},
	{200, 14, 2, Floor, 28},
	{1414104958, 21.45, 2, Floor, 303325513.49},
	{9857, 75.6, 1, Floor, 7451.8},
	{9857, 75.6, 0, Floor, 7451},
	{100, 10, 2, Floor, 10},
	{1055, 350, 2, Floor, 3692.5},
	{100, 1598, 2, Floor, 1598},
	{333, 8, 2, Floor, 26.64},
	{1992, .345, 2, Floor, 6.87},
	{65, .011, 3, Floor, 0.007},
	{65, .011, 4, Floor, 0.0071},

	{898, 56.7, 2, Ceil, 509.17},
	{898, 56.7, 3, Ceil, 509.166},
	{10975, 11, 2, Ceil, 1207.25},
	{10975, 11, 3, Ceil, 1207.25},
	{10, 27, 2, Ceil, 2.7},
	{6942, 99, 2, Ceil, 6872.58},
	{9999, 0, 2, Ceil, 0},
	{9999, 100, 2, Ceil, 9999},
	{45435, 69, 2, Ceil, 31350.15},
	{420, 42, 2, Ceil, 176.4},
	{89357, 85.4, 2, Ceil, 76310.88},
	{200, 14, 2, Ceil, 28},
	{1414104958, 21.45, 2, Ceil, 303325513.50},
	{9857, 75.6, 1, Ceil, 7451.9},
	{9857, 75.6, 0, Ceil, 7452},
	{100, 10, 2, Ceil, 10},
	{1055, 350, 2, Ceil, 3692.5},
	{100, 1598, 2, Ceil, 1598},
	{333, 8, 2, Ceil, 26.64},
	{1992, .345, 2, Ceil, 6.88},
	{65, .011, 3, Ceil, 0.008},
	{65, .011, 4, Ceil, 0.0072},

	{3, 50, 0, Bankers, 2},
	{5, 50, 0, Bankers, 2},
}

func TestGetPercentageFromInt(t *testing.T) {
	for _, v := range TestLargeNums {
		result := PercentageFromInt(v.Integer, 1, 2, Round)
		if result != v.Float2 {
			t.Error("Expected: ", v.Float2, "Got: ", result)
		}
	}
	for _, v := range intPercentageData {
		result := PercentageFromInt(v.amt, v.pct, v.fraction, v.round)
		if result != v.result {
			t.Error("Expected: ", v.result, "Got: ", result)
		}
	}
}

var floatPercentageData = []struct {
	amt      float64
	pct      float64
	fraction int
	round    round
	result   float64
}{
	{64.72, 10, 3, Round, 6.472},
	{64.72, 10, 2, Round, 6.47},
	{11.11, 13, 2, Round, 1.44},
	{11.11, 13, 4, Round, 1.4443},
	{9999.99, 100, 2, Round, 9999.99},
	{10000.85, 0, 2, Round, 0},
	{420.69, 42, 2, Round, 176.69},
	{1.25, 50, 1, Round, 0.6},
	{95545.194, 39.9, 2, Round, 38122.53},
	{95545.194, 39.9, 3, Round, 38122.532},
	{95545.194, 39.9, 4, Round, 38122.5324},
	{95545.194, 39.9, 5, Round, 38122.53241},
	{95545.194, 39.9, 6, Round, 38122.532406},
	{21.4, 1540, 2, Round, 329.56},
	{0.5, 1, 2, Round, 0.01},
	{0.5, 1, 3, Round, 0.005},
	{0, 42, 2, Round, 0},
	{134.2, 55.5, 0, Round, 74},
	{19.93, .045, 2, Round, .01},
	{19.93, .045, 4, Round, .009},
	{19.93, .045, 5, Round, .00897},
	{0.25, 40, 1, Round, .1},
	{0.25, 4, 2, Round, .01},
	{0.25, 4, 4, Round, .01},
	{-0.09, 1, 2, Round, 0},

	{64.72, 10, 3, Floor, 6.472},
	{64.72, 10, 2, Floor, 6.47},
	{11.11, 13, 2, Floor, 1.44},
	{11.11, 13, 4, Floor, 1.4443},
	{9999.99, 100, 2, Floor, 9999.99},
	{10000.85, 0, 2, Floor, 0},
	{420.69, 42, 2, Floor, 176.68},
	{1.25, 50, 1, Floor, 0.6},
	{95545.194, 39.9, 2, Floor, 38122.53},
	{95545.194, 39.9, 3, Floor, 38122.532},
	{95545.194, 39.9, 4, Floor, 38122.5324},
	{95545.194, 39.9, 5, Floor, 38122.53240},
	{95545.194, 39.9, 6, Floor, 38122.532406},
	{21.4, 1540, 2, Floor, 329.56},
	{0.5, 1, 2, Floor, 0},
	{0.5, 1, 3, Floor, 0.005},
	{0, 42, 2, Floor, 0},
	{134.2, 55.5, 0, Floor, 74},
	{19.93, .045, 2, Floor, 0},
	{19.93, .045, 4, Floor, .0089},
	{19.93, .045, 5, Floor, .00896},

	{64.72, 10, 3, Ceil, 6.472},
	{64.72, 10, 2, Ceil, 6.48},
	{11.11, 13, 2, Ceil, 1.45},
	{11.11, 13, 4, Ceil, 1.4443},
	{9999.99, 100, 2, Ceil, 9999.99},
	{10000.85, 0, 2, Ceil, 0},
	{420.69, 42, 2, Ceil, 176.69},
	{1.25, 50, 1, Ceil, 0.7},
	{95545.194, 39.9, 2, Ceil, 38122.54},
	{95545.194, 39.9, 3, Ceil, 38122.533},
	{95545.194, 39.9, 4, Ceil, 38122.5325},
	{95545.194, 39.9, 5, Ceil, 38122.53241},
	{95545.194, 39.9, 6, Ceil, 38122.532406},
	{21.4, 1540, 2, Ceil, 329.56},
	{0.5, 1, 2, Ceil, 0.01},
	{0.5, 1, 3, Ceil, 0.005},
	{0, 42, 2, Ceil, 0},
	{134.2, 55.5, 0, Ceil, 75},
	{19.93, .045, 2, Ceil, .01},
	{19.93, .045, 4, Ceil, .009},
	{19.93, .045, 5, Ceil, .00897},

	{1.5, 100, 0, Bankers, 2},
	{2.5, 100, 0, Bankers, 2},
	{1.535, 100, 2, Bankers, 1.54},
	{1.525, 100, 2, Bankers, 1.52},
	{0.5, 100, 0, Bankers, 0},
	{1.5, 100, 0, Bankers, 2},
	{0.4, 100, 0, Bankers, 0},
	{0.6, 100, 0, Bankers, 1},
	{1.4, 100, 0, Bankers, 1},
	{1.6, 100, 0, Bankers, 2},
	{23.5, 100, 0, Bankers, 24},
	{24.5, 100, 0, Bankers, 24},
	{-23.5, 100, 0, Bankers, -24},
	{-24.5, 100, 0, Bankers, -24},
	{1.534953, 100, 2, Bankers, 1.53},
	{1.53499999, 100, 2, Bankers, 1.53},
	{1.53499999999, 100, 2, Bankers, 1.53},
	{1.5299999, 100, 2, Bankers, 1.53},
	{1.5350, 100, 2, Bankers, 1.54},
	{1.53599, 100, 2, Bankers, 1.54},
	{3.2, 100, 0, Bankers, 3},
	{3.4, 100, 0, Bankers, 3},
	{3.5, 100, 0, Bankers, 4},
	{4.5, 100, 0, Bankers, 4},
	{5.5, 100, 0, Bankers, 6},
	{-7.5, 100, 0, Bankers, -8},
}

func TestGetPercentageFromFloat(t *testing.T) {
	for _, v := range floatPercentageData {
		result := PercentageFromFloat(v.amt, v.pct, v.fraction, v.round)
		if result != v.result {
			t.Error("Expected: ", v.result, "Got: ", result)
		}
	}
}

// randFloat generates a random floating point number to the given precision.
func randFloat(min, max float64, prec int) float64 {
	result := min + rand.Float64()*(max-min)
	return math.Round(result*math.Pow10(prec)) / math.Pow10(prec)
}

type roundingErrorTestInt struct {
	amt            int
	pct            float64
	fraction       int
	round          round
	maxMantissaLen int
}

type roundingErrorTestFloat struct {
	amt            float64
	pct            float64
	fraction       int
	round          round
	maxMantissaLen int
}

func TestRoundingError(t *testing.T) {
	var testNumsInt []roundingErrorTestInt
	var testNumsFloat []roundingErrorTestFloat

	for i := 1; i < 100000; i++ {
		testNum := roundingErrorTestInt{
			amt:            i,
			pct:            randFloat(1.0, 99.9, 4),
			fraction:       3,
			round:          Bankers,
			maxMantissaLen: 4,
		}

		testNumsInt = append(testNumsInt, testNum)
	}

	for i := float64(1); i < 100000.99; i += .01 {
		testNum := roundingErrorTestFloat{
			amt:            i,
			pct:            randFloat(1.0, 99.9, 4),
			fraction:       3,
			round:          Bankers,
			maxMantissaLen: 4,
		}

		testNumsFloat = append(testNumsFloat, testNum)
	}

	for _, v := range testNumsInt {
		result := PercentageFromInt(v.amt, v.pct, v.fraction, v.round)
		resultStr := strconv.FormatFloat(result, 'f', -1, 64)
		split := strings.Split(resultStr, ".")

		var resultMantLen int
		if len(split) == 2 {
			resultMantLen = len(split[1])
		}

		if resultMantLen > v.maxMantissaLen {
			t.Errorf("Expected mantissa length to be less than %v, got %v", v.maxMantissaLen, resultMantLen)
		}
	}

	for _, v := range testNumsFloat {
		result := PercentageFromFloat(v.amt, v.pct, v.fraction, v.round)
		resultStr := strconv.FormatFloat(result, 'f', -1, 64)
		split := strings.Split(resultStr, ".")

		var resultMantLen int
		if len(split) == 2 {
			resultMantLen = len(split[1])
		}

		if resultMantLen > v.maxMantissaLen {
			t.Errorf("Expected mantissa length to be less than %v, got %v", v.maxMantissaLen, resultMantLen)
		}
	}
}

func BenchmarkPercentageFromFloat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PercentageFromFloat(1.534935, 50.393, 2, Round)
	}
}
