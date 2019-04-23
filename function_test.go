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
