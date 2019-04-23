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
	{"AED", "AED"},
	{"ARS", "ARS"},
	{"aud", "AUD"},
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
	{"784", "784"},
	{"032", "032"},
	{"036", "036"},
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
	{"USD", Currency{Unit: "US Dollar", Alpha: "USD", Numeric: "840", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}},
	{"AED", Currency{Unit: "UAE Dirham", Alpha: "AED", Numeric: "784", Symbol: "\u0625\u002E\u062F", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: ","}},
	{"ARS", Currency{Unit: "Argentine Peso", Alpha: "ARS", Numeric: "032", Symbol: "\u0024", Exponent: 2, Decimal: ",", Separator: 3, Delimiter: "."}},
	{"AUD", Currency{Unit: "Australian Dollar", Alpha: "AUD", Numeric: "036", Symbol: "\u0024", Exponent: 2, Decimal: ".", Separator: 3, Delimiter: " "}},
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
	{"784", "AED"},
	{"032", "ARS"},
	{"036", "AUD"},
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
	{uint(0), 1, "0.0"},
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
