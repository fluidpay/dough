package dough

import "testing"

func TestInvalidCountryCode(t *testing.T) {
	ret := ValidatePostalCodeByCountyCode("ZZZ", "asdf")
	if ret != false {
		t.Errorf("Expected flase received true")
	}

}

func TestUSPostalCodes(t *testing.T) {
	if ValidatePostalCodeByCountyCode("US", "asdf") {
		t.Errorf("Expected flase received true")
	}
	if ValidatePostalCodeByCountyCode("US", "1234") {
		t.Errorf("Expected flase received true")
	}
	if !ValidatePostalCodeByCountyCode("US", "12345") {
		t.Errorf("Expected true received false")
	}
	if ValidatePostalCodeByCountyCode("US", "123456") {
		t.Errorf("Expected true received false")
	}
	if ValidatePostalCodeByCountyCode("US", "12345-6") {
		t.Errorf("Expected true received false")
	}
	if !ValidatePostalCodeByCountyCode("US", "12345-6789") {
		t.Errorf("Expected false received true")
	}
}

func TestCAPostalCodes(t *testing.T) {
	if ValidatePostalCodeByCountyCode("CA", "asdf") {
		t.Errorf("Expected flase received true")
	}
	if ValidatePostalCodeByCountyCode("CA", "12345") {
		t.Errorf("Expected flase received true")
	}
	if !ValidatePostalCodeByCountyCode("CA", "A1A 1A1") {
		t.Errorf("Expected true received false")
	}
	if !ValidatePostalCodeByCountyCode("CA", "X0E 0T0") {
		t.Errorf("Expected false received true")
	}
}

func TestGBPostalCodes(t *testing.T) {
	if ValidatePostalCodeByCountyCode("GB", "asdf") {
		t.Errorf("Expected flase received true")
	}
	if ValidatePostalCodeByCountyCode("GB", "12345") {
		t.Errorf("Expected flase received true")
	}
	if !ValidatePostalCodeByCountyCode("GB", "B1 1AA") {
		t.Errorf("Expected true received false")
	}
	if !ValidatePostalCodeByCountyCode("GB", "L1 0AA") {
		t.Errorf("Expected false received true")
	}
	if !ValidatePostalCodeByCountyCode("GB", "ZE1 0AA") {
		t.Errorf("Expected false received true")
	}
}
